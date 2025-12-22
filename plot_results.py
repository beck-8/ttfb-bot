import pandas as pd
import matplotlib.pyplot as plt
import seaborn as sns
import sys
import os

# Usage: python3 plot_results.py [csv_file]
# Example: python3 plot_results.py calibration_random.csv

def main():
    file_path = sys.argv[1] if len(sys.argv) > 1 else 'calibration_random.csv'
    
    if not os.path.exists(file_path):
        print(f"Error: File '{file_path}' not found.")
        sys.exit(1)

    print(f"Reading data from {file_path}...")
    try:
        df = pd.read_csv(file_path)
    except Exception as e:
        print(f"Error reading CSV: {e}")
        sys.exit(1)

    # Clean up column names / handle compatibility
    # Ensure we use 'AvgTTFB(ms)' if present, else 'TTFB(ms)'
    ttfb_col = 'AvgTTFB(ms)' if 'AvgTTFB(ms)' in df.columns else 'TTFB(ms)'
    speed_col = 'AvgSpeed(MB/s)' if 'AvgSpeed(MB/s)' in df.columns else 'Speed(MB/s)'
    p95_col = 'P95TTFB(ms)'
    
    # Filter for interesting rows (e.g., successful or at least attempted)
    if 'SuccessRate' in df.columns:
        # Plot everything that has a SuccessRate >= 0 (basically everything)
        # Maybe we want to highlight 0% success differently?
        df_plot = df.copy()
    else:
        # Backward compatibility
        if 'Success' in df.columns:
            # Convert 'true'/'false' strings to bool if needed
            df['Success'] = df['Success'].apply(lambda x: str(x).lower() == 'true')
            df['SuccessRate'] = df['Success'].apply(lambda x: 1.0 if x else 0.0)
            df_plot = df[df['Success'] == True].copy()
        else:
            df_plot = df.copy()

    if df_plot.empty:
        print("No data to plot.")
        sys.exit()

    # Get Context
    local_region = "Unknown"
    if 'LocalRegion' in df.columns and not df.empty:
        val = str(df['LocalRegion'].iloc[0])
        if val and val != "nan":
            local_region = val
            
    # Set global style
    sns.set_theme(style="whitegrid", context="talk") # 'talk' context makes fonts larger/readable
    
    # Create Figure: 3 Subplots (Latency, Speed, Success)
    # 2 rows, 2 columns? Or 3 rows?
    # Let's do 2 rows: Top=Latency+Speed, Bottom=Success+Table? 
    # Or just 3 vertical?
    # Let's stick to the 2x2 grid, it's efficient.
    
    fig, axes = plt.subplots(2, 2, figsize=(16, 12))
    plt.suptitle(f"Filecoin PDP Performance Benchmark\nRegion: {local_region}", fontsize=20, y=0.98)

    # --- 1. Latency (TTFB) ---
    ax_ttfb = axes[0, 0]
    # Sort by AvgTTFB ascending (fastest first)
    # Filter out 0ms if that means failure? Or keep them?
    # Usually failed tests might have 0ms. Let's filter for this specific plot.
    df_ttfb = df_plot[df_plot[ttfb_col] > 0].sort_values(ttfb_col)
    
    if not df_ttfb.empty:
        sns.barplot(data=df_ttfb, x='Name', y=ttfb_col, hue='Name', palette='viridis', ax=ax_ttfb, dodge=False)
        if ax_ttfb.legend_: ax_ttfb.legend_.remove()
        
        # Overlay P95
        if p95_col in df.columns:
            # We align markers with the bars
            sns.scatterplot(data=df_ttfb, x='Name', y=p95_col, color='black', marker='X', s=150, zorder=10, ax=ax_ttfb, label='P95')
            ax_ttfb.legend(loc='upper right', frameon=True)
            
        ax_ttfb.set_title('Latency (TTFB)', fontsize=16, fontweight='bold')
        ax_ttfb.set_ylabel('Time (ms)')
        ax_ttfb.set_xlabel('')
        ax_ttfb.tick_params(axis='x', rotation=45)
    else:
        ax_ttfb.text(0.5, 0.5, "No valid TTFB data", ha='center', va='center')

    # --- 2. Speed (Throughput) ---
    ax_speed = axes[0, 1]
    df_speed = df_plot[df_plot[speed_col] > 0].sort_values(speed_col, ascending=False)
    
    if not df_speed.empty:
        sns.barplot(data=df_speed, x='Name', y=speed_col, hue='Name', palette='magma', ax=ax_speed, dodge=False)
        if ax_speed.legend_: ax_speed.legend_.remove()
        ax_speed.set_title('Throughput (Download Speed)', fontsize=16, fontweight='bold')
        ax_speed.set_ylabel('Speed (MB/s)')
        ax_speed.set_xlabel('')
        ax_speed.tick_params(axis='x', rotation=45)
    else:
        ax_speed.text(0.5, 0.5, "No valid Speed data", ha='center', va='center')

    # --- 3. Reliability (Success Rate) ---
    ax_rate = axes[1, 0]
    # Show all providers here
    df_rate = df_plot.sort_values('SuccessRate', ascending=False)
    
    if not df_rate.empty:
        sns.barplot(data=df_rate, x='Name', y='SuccessRate', hue='Name', palette='RdYlGn', ax=ax_rate, dodge=False)
        if ax_rate.legend_: ax_rate.legend_.remove()
        ax_rate.set_title('Reliability (Success Rate)', fontsize=16, fontweight='bold')
        ax_rate.set_ylabel('Rate (0.0 - 1.0)')
        ax_rate.set_xlabel('')
        ax_rate.set_ylim(0, 1.1) 
        ax_rate.tick_params(axis='x', rotation=45)
    else:
        ax_rate.text(0.5, 0.5, "No Success Rate data", ha='center', va='center')
        
    # --- 4. Summary Text / Stats ---
    ax_info = axes[1, 1]
    ax_info.axis('off')
    
    # Let's write some summary stats as text
    total_providers = len(df)
    total_success = len(df_plot[df_plot['SuccessRate'] > 0.9]) # Arbitrary threshold for "Good"
    fastest_provider = df_ttfb.iloc[0]['Name'] if not df_ttfb.empty else "N/A"
    fastest_ttfb = df_ttfb.iloc[0][ttfb_col] if not df_ttfb.empty else 0
    
    summary_text = (
        f"Summary Stats:\n\n"
        f"Total Providers Tested: {total_providers}\n"
        f"High Reliability (>90%): {total_success}\n\n"
        f"Fastest Responder: {fastest_provider} ({fastest_ttfb} ms)\n"
    )
    
    if not df_speed.empty:
        highest_speed_prov = df_speed.iloc[0]['Name']
        highest_speed_val = df_speed.iloc[0][speed_col]
        summary_text += f"Highest Speed: {highest_speed_prov} ({highest_speed_val:.2f} MB/s)\n"
        
    ax_info.text(0.1, 0.5, summary_text, fontsize=14, linespacing=1.5, verticalalignment='center')


    plt.tight_layout()
    plt.subplots_adjust(hspace=0.6, top=0.90)

    output_file = 'performance_summary_polished.png'
    plt.savefig(output_file, dpi=150) # Higher DPI
    print(f"Success! Plot saved to {output_file}")
    
    if sys.platform == 'darwin':
        os.system(f'open {output_file}')

if __name__ == "__main__":
    main()

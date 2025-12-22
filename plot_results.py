import pandas as pd
import matplotlib.pyplot as plt
import seaborn as sns
import sys
import os

# Usage: python3 plot_results.py [csv_file]

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

    # Compat: Map new column names to generic variables
    ttfb_col = 'AvgTTFB(ms)' if 'AvgTTFB(ms)' in df.columns else 'TTFB(ms)'
    speed_col = 'AvgSpeed(MB/s)' if 'AvgSpeed(MB/s)' in df.columns else 'Speed(MB/s)'
    p95_col = 'P95TTFB(ms)'
    
    # Filter Success
    if 'SuccessRate' in df.columns:
        df_plot = df[df['SuccessRate'] > 0].copy()
    elif 'Success' in df.columns:
        # Check for string 'true' or boolean True
        df['Success'] = df['Success'].apply(lambda x: str(x).lower() == 'true')
        df_plot = df[df['Success'] == True].copy()
    else:
        df_plot = df.copy()

    if df_plot.empty:
        print("No successful tests to plot.")
        sys.exit()

    # Get Context
    local_region = ""
    if 'LocalRegion' in df.columns and not df.empty:
        val = str(df['LocalRegion'].iloc[0])
        if val and val != "nan":
            local_region = f" (Tester: {val})"

    # Style: Return to simple clean style
    sns.set_theme(style="whitegrid")
    
    # 1 Row, 2 Columns (Classic Look)
    plt.figure(figsize=(14, 6))
    plt.suptitle(f"Filecoin PDP Performance{local_region}", fontsize=14)

    # --- Plot 1: Latency (TTFB) ---
    plt.subplot(1, 2, 1)
    # Sort for readability
    df_sorted = df_plot.sort_values(ttfb_col)
    
    ax1 = sns.barplot(data=df_sorted, x='Name', y=ttfb_col, hue='Name', palette='viridis', dodge=False)
    if ax1.legend_: ax1.legend_.remove()
    
    # Optional: Subtle P95 marker overlay
    if p95_col in df.columns:
        sns.scatterplot(data=df_sorted, x='Name', y=p95_col, color='black', marker='_', s=100, label='P95', zorder=10)
        plt.legend()
    
    plt.title(f'Latency ({ttfb_col})\nLower is Better')
    plt.ylabel('Time (ms)')
    plt.xlabel('')
    plt.xticks(rotation=45, ha='right')

    # --- Plot 2: Speed ---
    plt.subplot(1, 2, 2)
    df_sorted_speed = df_plot.sort_values(speed_col, ascending=False)
    
    ax2 = sns.barplot(data=df_sorted_speed, x='Name', y=speed_col, hue='Name', palette='magma', dodge=False)
    if ax2.legend_: ax2.legend_.remove()
    
    plt.title(f'Throughput ({speed_col})\nHigher is Better')
    plt.ylabel('Speed (MB/s)')
    plt.xlabel('')
    plt.xticks(rotation=45, ha='right')

    plt.tight_layout()
    plt.subplots_adjust(top=0.85)

    output_file = 'performance_summary_clean.png'
    plt.savefig(output_file)
    print(f"Success! Plot saved to {output_file}")
    
    if sys.platform == 'darwin':
        os.system(f'open {output_file}')

if __name__ == "__main__":
    main()

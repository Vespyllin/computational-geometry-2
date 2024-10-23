import pandas as pd
import matplotlib.pyplot as plt
import os

# Get all CSV files in the current directory
csv_files = [f for f in os.listdir('.') if f.endswith('.csv')]

for file in csv_files:
    # Load the CSV file
    df = pd.read_csv(file)

    # Parse points from the columns
    main_points = df['Main Points'].dropna().str.split(',', expand=True).astype(float)
    hull_points = df['Convex Hull Points'].dropna().str.split(',', expand=True).astype(float)

    # Create scatter plot for main points
    plt.scatter(main_points[0], main_points[1], color='blue', label='Main Points')

    # Create line plot for convex hull
    plt.plot(hull_points[0], hull_points[1], color='red', label='Convex Hull')

    # Highlight points in the convex hull
    plt.scatter(hull_points[0], hull_points[1], color='red')

    # Add labels and title
    plt.xlabel('X')
    plt.ylabel('Y')
    plt.title(f'Convex Hull Visualization - {file}')
    plt.legend()
    plt.grid(True)

    # Save the plot as a PNG file based on CSV file name
    plot_filename = file.replace('.csv', '.png')
    plt.savefig(plot_filename)
    plt.close()  # Close the figure to avoid overlap in the next loop

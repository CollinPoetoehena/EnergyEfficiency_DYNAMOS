import pandas as pd
from sklearn.cluster import DBSCAN
import argparse

import utils

# Detect anomalies using DBSCAN
def execute_DBSCAN_AD_algorithm(df: pd.DataFrame, exp_dirs):
    column_to_use = 'total_energy_difference'
    X = df[[column_to_use]].values

    # DBSCAN parameters, increase eps to detect fewer anomalies
    dbscan = DBSCAN(eps=47, min_samples=3)
    labels = dbscan.fit_predict(X)

    # Anomalies are labeled as -1
    anomalies = df[labels == -1]
    print("Anomalies detected:")
    print(anomalies)

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description="Run energy efficiency experiment")
    parser.add_argument("archetype", type=str, choices=["ComputeToData", "DataThroughTTP"], 
                        help="Archetype to detect anomalies from.")
    parser.add_argument("prefix", type=str, help="Prefix of the experiment folders")
    args = parser.parse_args()

    # Load the data
    df, exp_dirs = utils.load_experiment_results(args.prefix, args.archetype)

    if df.empty:
        print("No data loaded. Exiting.")
    else:
        # Call your new DBSCAN-based function
        execute_DBSCAN_AD_algorithm(df, exp_dirs)

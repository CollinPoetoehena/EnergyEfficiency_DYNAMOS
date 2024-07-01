from utils import query_prometheus, save_energy_data_to_file
from constants import RELEVANT_PROMETHEUS_CONTAINER_METRICS, RELEVANT_KEPLER_CONTAINER_METRICS, PROMETHEUS_QUERIES

def main():
    # Calculate energy consumption
    energy_data = calculate_energy_consumption()
    
    # Add the relevant metrics data
    energy_data.update(get_relevant_metrics())

    # Save the data to a file
    save_energy_data_to_file(energy_data, 'energy_metrics')


# Function to calculate energy consumption and save to a file
def calculate_energy_consumption():
    output_data = {}
    # Query Prometheus to get the energy consumption
    result = query_prometheus(PROMETHEUS_QUERIES['ENERGY_CONSUMPTION'])
    # Check if the query was successful
    if result['status'] == 'success':
        # Add the data to the output with the key 'energy_consumption'
        output_data['energy_consumption'] = result['data']['result']
    else:
        raise Exception(f"Failed to retrieve data: {result['error']}")

    # Return the data
    return output_data

# Function to query relevant metrics
def get_relevant_metrics():
    output_data = {}

    # Query and filter RELEVANT_KEPLER_CONTAINER_METRICS
    for metric in RELEVANT_KEPLER_CONTAINER_METRICS:
        # Query Prometheus to get the metric
        result = query_prometheus(metric)
        # Check if the query was successful
        if result['status'] == 'success':
            # Filter out metrics with 0 values (not measured in Kepler)
            # The value will look like this: {metric: [{metric: {key: value}..., value: [timestamp, value]}]
            filtered_data = [
                item for item in result['data']['result'] if float(item['value'][1]) != 0
            ]
            if filtered_data:
                output_data[metric] = filtered_data
        else:
            print(f"Failed to retrieve data for metric {metric}: {result['error']}")

    # Query and add RELEVANT_PROMETHEUS_CONTAINER_METRICS
    for metric in RELEVANT_PROMETHEUS_CONTAINER_METRICS:
        # Query Prometheus to get the metric
        result = query_prometheus(metric)
        # Check if the query was successful
        if result['status'] == 'success':
            # Add the data to the output with the key {metric}
            output_data[metric] = result['data']['result']
        else:
            print(f"Failed to retrieve data for metric {metric}: {result['error']}")

    # Return the data
    return output_data

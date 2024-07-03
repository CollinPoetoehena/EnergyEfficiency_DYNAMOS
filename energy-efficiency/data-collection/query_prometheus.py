from requests import Response
import requests
from constants import PROMETHEUS_URL, CONTAINERS, KEPLER_CONTAINER_NAME_LABEL, CADVISOR_CONTAINER_NAME_LABEL, METRIC_STEP
from utils import convert_float_time_to_string

def exec_query(query: str, start_time: float, end_time: float) -> dict:
    """
    TODO: Add docstring. Function to query Prometheus.
    """
    print(f"Querying Prometheus with query: {query}")
    print(f"Start time: {convert_float_time_to_string(start_time)}, End time: {convert_float_time_to_string(end_time)}")

    # Query Prometheus with a time range
    response = requests.get(
        f"{PROMETHEUS_URL}/api/v1/query_range",
        params={
            "query": query,
            "start": start_time,
            "end": end_time,
            "step": f"{METRIC_STEP}s",
        },
    )

    # print(f"Response: {response.json()["data"]}")
    # Use the helper function to return the parsed data from the response
    return parse_prometheus_response(response)


def parse_prometheus_response(response: Response) -> dict:
    # If the query was successful, return the results
    if response.status_code == 200:
        # Initialize a dictoinary to store the data
        data = {}
        # Get the result from the response
        results = response.json()["data"]["result"]
        # print(f"Results: {results}")
        
        # Convert results to string and print the first 500 characters
        results_str = str(results)
        print(f"Results: {results_str[:500]}")

        # Loop through the results
        for result in results:
            # Initialize container_name as None
            container_name = None

            # Check if the result contains specific keys (required to identify the container used)
            if not all(
                k not in result["metric"].keys() for k in [KEPLER_CONTAINER_NAME_LABEL, CADVISOR_CONTAINER_NAME_LABEL]
            ):
                return

            print(result["metric"].keys())

            # Get the container name from the result
            if KEPLER_CONTAINER_NAME_LABEL in result["metric"]:
                container_name = result["metric"][KEPLER_CONTAINER_NAME_LABEL]
            elif CADVISOR_CONTAINER_NAME_LABEL in result["metric"]:
                container_name = result["metric"][CADVISOR_CONTAINER_NAME_LABEL]
        
            # Only add the data if the container name is in the containers list
            if container_name in CONTAINERS:
                data[container_name] = result["values"]

        # Return the data
        return data
    else:
        raise Exception(
            f"Query failed with status code {response.status_code}: {response.content}")
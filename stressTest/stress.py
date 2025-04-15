import asyncio
import aiohttp
import time

# Configuration
URLS = {
    "Gin (GO)": "http://localhost:8080/ping",
    "FastAPI (Python)": "http://localhost:8079/ping"
}

NUM_REQUESTS = 5000       # Total number of requests
CONCURRENCY_LIMIT = 1000  # Maximum concurrent requests
REQUEST_TIMEOUT = 30.0    # Timeout in seconds

HEADERS = {
    "accept": "application/json",
    "user-agent": "Mozilla/5.0"
}

async def fetch(session, url):
    """Send a single GET request."""
    try:
        async with session.get(url, headers=HEADERS, timeout=REQUEST_TIMEOUT) as response:
            return await response.text()
    except asyncio.TimeoutError:
        return "Timeout"
    except Exception as e:
        return f"Error: {str(e)}"


async def stress_test(url, num_requests, concurrency_limit):
    """Perform a stress test on the given URL."""
    connector = aiohttp.TCPConnector(limit=concurrency_limit)
    async with aiohttp.ClientSession(connector=connector) as session:
        tasks = [fetch(session, url) for _ in range(num_requests)]
        start_time = time.time()
        responses = await asyncio.gather(*tasks)
        end_time = time.time()
        
        # Count successful vs failed responses
        timeouts = responses.count("Timeout")
        errors = sum(1 for r in responses if r.startswith("Error:"))
        successful = len(responses) - timeouts - errors
        
        return {
            "total": len(responses),
            "successful": successful,
            "timeouts": timeouts,
            "errors": errors,
            "duration": end_time - start_time
        }


async def main():
    """Run stress tests for both servers."""
    for name, url in URLS.items():
        print(f"Starting stress test for {name}...")
        results = await stress_test(url, NUM_REQUESTS, CONCURRENCY_LIMIT)
        print(f"{name} Results:")
        print(f"  Total Requests: {results['total']}")
        print(f"  Successful Responses: {results['successful']}")
        print(f"  Timeouts: {results['timeouts']}")
        print(f"  Errors: {results['errors']}")
        print(f"  Total Time: {results['duration']:.2f} seconds")
        print(f"  Requests per Second: {results['total'] / results['duration']:.2f} RPS")
        print("-" * 40)


if __name__ == "__main__":
    try:
        asyncio.run(main())
    except Exception as e:
        print(f"An error occurred: {e}")
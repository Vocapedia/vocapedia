const cache = new Map();

export async function useFetch(
    endpoint,
    {
        method = "GET",
        body = null,
        headers = {},
        timeout = 10000,
        cacheTTL = 60000,
        retry = 3,
        retryDelay = 1000
    } = {}
) {
    const baseURL = window.location.hostname === "localhost"
        ? "http://localhost:3000/api/v1"
        : "https://vocapedia.space/api/v1"
    
        const token = localStorage.getItem("token");

    const defaultHeaders = {
        "Content-Type": "application/json",
        ...(token ? { Authorization: `Bearer ${token}` } : {}),
        ...headers,
    };

    const config = {
        method,
        headers: defaultHeaders,
    };

    if (body) {
        config.body = JSON.stringify(body);
    }

    const url = `${baseURL}${endpoint}`;

    if (method === "GET" && cache.has(url)) {
        const { data, timestamp } = cache.get(url);
        if (Date.now() - timestamp < cacheTTL) {
            console.log(`Cache'den alındı: ${url}`);
            return data;
        } else {
            cache.delete(url);
        }
    }

    let attempt = 0;

    while (attempt < retry) {
        attempt++;

        const controller = new AbortController();
        const timeoutId = setTimeout(() => controller.abort(), timeout);
        config.signal = controller.signal;

        try {
            const response = await fetch(url, config);
            clearTimeout(timeoutId);

            if (!response.ok) {
                throw new Error(`HTTP error! Status: ${response.status}`);
            }

            const data = await response.json();

            if (method === "GET") {
                cache.set(url, { data, timestamp: Date.now() });
            }

            return data;
        } catch (error) {
            console.error(`Fetch attempt ${attempt} failed:`, error);

            if (attempt >= retry) {
                throw error;
            }

            await new Promise(resolve => setTimeout(resolve, retryDelay)); // Tekrar denemeden önce bekle
        }
    }
}

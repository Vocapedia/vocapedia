const cache = new Map();
const mode = import.meta.env.VITE_MODE;
const built = import.meta.env.VITE_BUILT;
import _search_list from "@/fake/search_list.json";
import _trends_list from '@/fake/trends_list.json'
import _chapter from "@/fake/chapter.json";

export async function useFetch(
    endpoint,
    {
        method = "GET",
        body = null,
        headers = {},
        timeout = 10000,
        cacheTTL = 20000,
        retry = 3,
        retryDelay = 1000
    } = {}
) {
    if (mode !== "development") {
        return await Fetch(endpoint, { method, body, headers, timeout, cacheTTL, retry, retryDelay });
    }
    if (built == 1) {
        return await Fetch(endpoint, { method, body, headers, timeout, cacheTTL, retry, retryDelay });
    }
    return await Mock(endpoint, method);
}

const snowflakeRegex = /^\d{10,19}$/;

async function Mock(endpoint, method = "GET") {
    await new Promise(resolve => setTimeout(resolve, 500));

    var term = endpoint.split("?")[0];
    term = term.split("/").map((item) => (snowflakeRegex.test(item) ? ":id" : item)).join("/");
    if (method === "GET") {
        switch (term) {
            case "/public/chapters/search":
                return _search_list;
            case "/chapters/favorite":
                return _search_list;
            case "/public/chapters/trends":
                return _trends_list;
            case "/public/chapters/:id":
                return _chapter;
            case "/public/chapters/user":
                return _search_list;
            default:
                return "{}";
        }
    }
}

async function Fetch(endpoint, { method, body, headers, timeout, cacheTTL, retry, retryDelay } = {}) {
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
            if (response.status == 401) {
                window.location.href = "/login"
            }
            if (!response.ok) {
                throw await response.json();
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

            await new Promise(resolve => setTimeout(resolve, retryDelay));
        }
    }

}
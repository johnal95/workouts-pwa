import type { StaticParse, TSchema } from "typebox";
import Value from "typebox/value";

async function fetchJSON<T extends TSchema>(Schema: T, url: string, init: RequestInit = {}): Promise<StaticParse<T>> {
    const headers = new Headers(init.headers);
    headers.set("accept", "application/json");

    const response = await fetch(url, { ...init, headers });

    if (!response.ok) {
        const cause = await response.text().catch(() => "unknown cause");
        throw new Error(`${init?.method ?? "GET"} ${url} failed with status ${response.status}: ${cause}`);
    }

    const responseBody = await response.json();

    return Value.Parse(Schema, responseBody);
}

export { fetchJSON };

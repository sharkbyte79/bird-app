import { birdObservation } from "../../types/shared.types.ts";

const BASE_URL: string = "http://localhost:8080";

export interface birdupRequest {
    method?: "GET" | "POST";
    body?: any;
}

export interface birdupResponse<T> {
    status: number;
    data: T;
}

export default async function birdupFetch<T = any>(endpoint: string, {
    method,
    body
}: birdupRequest): Promise<birdupResponse<T>> {
    // Join requested endpoint with base url
    const url: string = BASE_URL + endpoint;
    const res: Response = await fetch(url, {
        method,
        body
    });

    if (!res.ok) {
        throw new Error();
    }

    const data = await res.json();
    return {
        status: res.status,
        data: data as T
    }
}

export async function getObservations(regionCode: string, notable: boolean): Promise<birdupResponse<birdObservation[]>> {
    const endpoint: string = `/observations/${regionCode}${notable ? "/notable" : ""}`;

    return await birdupFetch(endpoint, {
        method: "GET",
    });
}
    

import { dev } from '$app/environment';
import { readable } from "svelte/store";

let _base: string;

if (dev) {
    _base = "http://localhost:7171/api"
} else {
    _base = "/api"
}

export const apiBase = readable(_base);
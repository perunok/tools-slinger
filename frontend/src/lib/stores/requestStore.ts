import { writable } from 'svelte/store';

export const requestStore = writable({
  request: null,
  body: null
});
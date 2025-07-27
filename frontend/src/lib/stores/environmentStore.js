// src/lib/stores/environmentStore.js
import { writable } from 'svelte/store';

// Sample initial data for environments
const initialEnvironments = [
  {
    id: 'env1',
    name: 'Development',
    variables: [
      { key: 'API_URL', value: 'https://dev.api.example.com' },
      { key: 'API_KEY', value: 'dev-key-123' }
    ]
  },
  {
    id: 'env2',
    name: 'Production',
    variables: [
      { key: 'API_URL', value: 'https://api.example.com' },
      { key: 'API_KEY', value: 'prod-key-456' }
    ]
  }
];

export const environmentStore = writable(initialEnvironments);
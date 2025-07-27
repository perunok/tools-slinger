<script>
  import { historyStore } from "../../lib/stores/historyStore";
  export let isDarkMode;
</script>

<div class="flex-1 overflow-y-auto">
  {#each $historyStore as request}
    <div
      class="p-2 hover:{isDarkMode
        ? 'bg-gray-600'
        : 'bg-gray-100'} rounded cursor-pointer"
    >
      <span
        class="font-medium {request.method === 'GET'
          ? isDarkMode
            ? 'text-green-400'
            : 'text-green-600'
          : isDarkMode
            ? 'text-blue-400'
            : 'text-blue-600'}"
      >
        {request.method}
      </span>
      <span
        class="text-sm {isDarkMode ? 'text-gray-400' : 'text-gray-500'} ml-2"
        >{request.url}</span
      >
    </div>
  {/each}
</div>

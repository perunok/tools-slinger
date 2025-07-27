<script>
  import { requestStore } from "../../lib/stores/requestStore";
  export let isDarkMode;
</script>

<div
  class="{isDarkMode
    ? 'dark:bg-gray-800 dark:text-white dark:border-gray-600'
    : 'bg-white text-black border-gray-200'} p-4 rounded shadow-md border"
>
  <h3 class="font-medium mb-2">Response</h3>
  {#if $requestStore.body}
    <pre
      class="{isDarkMode
        ? 'dark:bg-gray-700 dark:text-white'
        : 'bg-gray-100 text-black'} p-2 rounded">{$requestStore.body}</pre>
  {:else}
    <p class={isDarkMode ? "dark:text-gray-400" : "text-gray-500"}>
      No response yet. Send a request to see the response.
    </p>
  {/if}
</div>

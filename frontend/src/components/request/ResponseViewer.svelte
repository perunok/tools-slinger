<script>
  import { requestStore } from '../../lib/stores/requestStore';
</script>

<div class="bg-white p-4 rounded shadow-md">
  <h3 class="font-medium mb-2">Response</h3>
  {#if $requestStore.body}
    <pre class="bg-gray-100 p-2 rounded">{$requestStore.body}</pre>
  {:else}
    <p class="text-gray-500">No response yet. Send a request to see the response.</p>
  {/if}
</div>
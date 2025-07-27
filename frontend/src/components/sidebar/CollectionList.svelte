<script>
  import { collectionStore } from "../../lib/stores/collectionStore";
  export let isDarkMode;
</script>

<div class="flex-1 overflow-y-auto">
  {#each $collectionStore as collection}
    <div
      class="p-2 hover:{isDarkMode
        ? 'bg-gray-600'
        : 'bg-gray-100'} rounded cursor-pointer"
    >
      <span class="font-medium {isDarkMode ? 'text-white' : 'text-black'}"
        >{collection.name}</span
      >
      <span class="text-sm {isDarkMode ? 'text-gray-400' : 'text-gray-500'}"
        >({collection.requests.length} requests)</span
      >
    </div>
  {/each}
</div>

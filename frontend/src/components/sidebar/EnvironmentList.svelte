
<script>
  import { environmentStore } from '../../lib/stores/environmentStore';
  export let isDarkMode;
</script>

<div class="flex-1 overflow-y-auto">
  {#each $environmentStore as environment}
    <div class="p-2 hover:{isDarkMode ? 'bg-gray-600' : 'bg-gray-100'} rounded cursor-pointer">
      <span class="font-medium {isDarkMode ? 'text-white' : 'text-black'}">{environment.name}</span>
      <span class="text-sm {isDarkMode ? 'text-gray-400' : 'text-gray-500'}">({environment.variables.length} variables)</span>
    </div>
  {/each}
</div>

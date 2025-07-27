<script>
  export let isDarkMode;
  let tabs = [{ id: 1, name: "Request 1" }];
  let activeTab = 1;

  function addTab() {
    const newId = tabs.length > 0 ? Math.max(...tabs.map((t) => t.id)) + 1 : 1;
    tabs = [...tabs, { id: newId, name: `Request ${newId}` }];
    activeTab = newId;
  }

  function removeTab(id, event) {
    event.stopPropagation();
    tabs = tabs.filter((tab) => tab.id !== id);
    if (activeTab === id && tabs.length > 0) {
      activeTab = tabs[tabs.length - 1].id;
    } else if (tabs.length === 0) {
      activeTab = null;
    }
  }
</script>

<div class="flex items-center mb-4">
  {#each tabs as tab}
    <button
      class="px-4 py-2 mr-2 rounded-t flex items-center gap-2 {activeTab ===
      tab.id
        ? isDarkMode
          ? 'dark:bg-gray-800 dark:border-gray-600 dark:text-white'
          : 'bg-white border-t border-l border-r text-black'
        : isDarkMode
          ? 'dark:bg-gray-700 dark:text-gray-300'
          : 'bg-gray-200 text-gray-600'}"
      on:click={() => (activeTab = tab.id)}
    >
      <span>{tab.name}</span>
      <button
        class={isDarkMode
          ? "text-red-400 hover:text-red-300"
          : "text-red-500 hover:text-red-700"}
        on:click={(e) => removeTab(tab.id, e)}
        aria-label={`Close ${tab.name}`}
      >
        &times;
      </button>
    </button>
  {/each}
  <button
    class="{isDarkMode
      ? 'dark:text-blue-400 dark:hover:text-blue-300'
      : 'text-blue-600 hover:text-blue-700'} text-xl font-bold ml-2"
    on:click={addTab}
    aria-label="Add new tab"
  >
    +
  </button>
</div>

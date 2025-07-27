<script>
  import CollectionList from "./CollectionList.svelte";
  import HistoryList from "./HistoryList.svelte";
  import EnvironmentList from "./EnvironmentList.svelte";
  export let isSidebarOpen;
  export let toggleSidebar;
  export let isDarkMode;

  let secondarySidebarOpen = false;
  let secondarySidebarContent = null;

  function toggleSecondarySidebar(content) {
    if (secondarySidebarContent === content && secondarySidebarOpen) {
      secondarySidebarOpen = false;
      secondarySidebarContent = null;
    } else {
      secondarySidebarContent = content;
      secondarySidebarOpen = true;
    }
  }
</script>

<div class="flex h-full">
  <!-- Main Sidebar Content -->
  <div class="flex flex-col p-4 h-full">
    <div class="flex justify-between items-center mb-4">
      <h1 class="text-xl font-bold {isDarkMode ? 'text-white' : 'text-black'}">
        Slinger
      </h1>
      <button
        class="md:hidden {isDarkMode ? 'text-gray-300' : 'text-gray-600'}"
        on:click={toggleSidebar}
      >
        ✕
      </button>
    </div>
    <nav class="mb-4">
      <button
        class="w-full text-left py-2 px-4 rounded {secondarySidebarContent ===
        'collections'
          ? isDarkMode
            ? 'bg-blue-700 text-blue-200'
            : 'bg-blue-100 text-blue-600'
          : isDarkMode
            ? 'text-gray-300'
            : 'text-gray-600'}"
        on:click={() => toggleSecondarySidebar("collections")}
      >
        Collections
      </button>
      <button
        class="w-full text-left py-2 px-4 rounded {secondarySidebarContent ===
        'history'
          ? isDarkMode
            ? 'bg-blue-700 text-blue-200'
            : 'bg-blue-100 text-blue-600'
          : isDarkMode
            ? 'text-gray-300'
            : 'text-gray-600'}"
        on:click={() => toggleSecondarySidebar("history")}
      >
        History
      </button>
      <button
        class="w-full text-left py-2 px-4 rounded {secondarySidebarContent ===
        'environments'
          ? isDarkMode
            ? 'bg-blue-700 text-blue-200'
            : 'bg-blue-100 text-blue-600'
          : isDarkMode
            ? 'text-gray-300'
            : 'text-gray-600'}"
        on:click={() => toggleSecondarySidebar("environments")}
      >
        Environments
      </button>
    </nav>
  </div>

  <!-- Secondary Sidebar -->
  {#if secondarySidebarOpen}
    <div
      class="flex flex-col w-64 p-4 {isDarkMode
        ? 'dark:bg-gray-700 dark:border-gray-600'
        : 'bg-white border-gray-200'} border-r shadow-lg h-full transition-all duration-300 absolute {isSidebarOpen
        ? 'left-64 md:left-64'
        : 'left-0 md:left-64'} z-10"
    >
      {#if secondarySidebarContent === "collections"}
        <div class="flex justify-between items-center mb-4">
          <h2
            class="text-lg font-semibold {isDarkMode
              ? 'text-white'
              : 'text-black'}"
          >
            Collections
          </h2>
          <button
            class={isDarkMode ? "text-gray-300" : "text-gray-600"}
            on:click={() => (secondarySidebarOpen = false)}>✕</button
          >
        </div>
        <CollectionList {isDarkMode} />
      {:else if secondarySidebarContent === "history"}
        <div class="flex justify-between items-center mb-4">
          <h2
            class="text-lg font-semibold {isDarkMode
              ? 'text-white'
              : 'text-black'}"
          >
            History
          </h2>
          <button
            class={isDarkMode ? "text-gray-300" : "text-gray-600"}
            on:click={() => (secondarySidebarOpen = false)}>✕</button
          >
        </div>
        <HistoryList {isDarkMode} />
      {:else if secondarySidebarContent === "environments"}
        <div class="flex justify-between items-center mb-4">
          <h2
            class="text-lg font-semibold {isDarkMode
              ? 'text-white'
              : 'text-black'}"
          >
            Environments
          </h2>
          <button
            class={isDarkMode ? "text-gray-300" : "text-gray-600"}
            on:click={() => (secondarySidebarOpen = false)}>✕</button
          >
        </div>
        <EnvironmentList {isDarkMode} />
      {/if}
    </div>
  {/if}
</div>

<style>
  @media (max-width: 768px) {
    .absolute {
      left: 0 !important;
    }
  }
</style>

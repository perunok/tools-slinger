<script>
  import Sidebar from "./components/sidebar/Sidebar.svelte";
  import Workspace from "./components/workspace/Workspace.svelte";
  import { onMount } from "svelte";

  let isSidebarOpen = true; // Sidebar visible by default
  let isDarkMode = false; // Default to light mode

  function toggleSidebar() {
    isSidebarOpen = !isSidebarOpen;
  }

  function toggleTheme() {
    isDarkMode = !isDarkMode;
    document.documentElement.classList.toggle("dark", isDarkMode);
  }

  // Persist sidebar and theme state
  onMount(() => {
    const savedSidebar = localStorage.getItem("isSidebarOpen");
    if (savedSidebar) isSidebarOpen = JSON.parse(savedSidebar);
    const savedTheme = localStorage.getItem("isDarkMode");
    if (savedTheme) {
      isDarkMode = JSON.parse(savedTheme);
      document.documentElement.classList.toggle("dark", isDarkMode);
    }
  });
  $: {
    localStorage.setItem("isSidebarOpen", JSON.stringify(isSidebarOpen));
    localStorage.setItem("isDarkMode", JSON.stringify(isDarkMode));
  }
</script>

<div class="flex h-screen {isDarkMode ? 'dark bg-gray-900' : 'bg-gray-100'}">
  <!-- Theme Toggle -->
  <button
    class="fixed top-4 right-4 z-30 {isDarkMode
      ? 'text-gray-300'
      : 'text-gray-600'}"
    on:click={toggleTheme}
    title={isDarkMode ? "Switch to Light Mode" : "Switch to Dark Mode"}
  >
    {#if isDarkMode}
      ☀️ <!-- Sun icon for light mode -->
    {:else}
      🌙 <!-- Moon icon for dark mode -->
    {/if}
  </button>

  <!-- Sidebar -->
  <div
    class="{isSidebarOpen ? 'w-64' : 'w-0'} {isDarkMode
      ? 'dark:bg-gray-800 dark:border-gray-700'
      : 'bg-white border-gray-200'} shadow-md h-full transition-all duration-300 md:w-64 overflow-hidden border-r"
  >
    <Sidebar {isSidebarOpen} {toggleSidebar} {isDarkMode} />
  </div>

  <!-- Hamburger Menu for Mobile -->
  {#if !isSidebarOpen}
    <button
      class="fixed top-4 left-4 md:hidden {isDarkMode
        ? 'text-gray-300'
        : 'text-gray-600'} z-20"
      on:click={toggleSidebar}
    >
      ☰
    </button>
  {/if}

 
  <div class="flex-1 flex flex-col">
    <Workspace {isDarkMode} />
  </div>
</div>

<style>
  @media (max-width: 768px) {
    .w-0 {
      width: 0;
      padding: 0;
    }
  }
</style>

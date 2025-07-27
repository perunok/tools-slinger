<script>
  import TabManager from './TabManager.svelte';
  import RequestBuilder from '../request/RequestBuilder.svelte';
  import ResponseViewer from '../request/ResponseViewer.svelte';
  export let isDarkMode;
</script>

<div class="{isDarkMode ? 'dark:bg-gray-800 dark:text-white' : 'bg-white text-black'}  rounded flex-1 flex flex-col p-4">
  <TabManager />
  <RequestBuilder />
  <ResponseViewer />
</div>
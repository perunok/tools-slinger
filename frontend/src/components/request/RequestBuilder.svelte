<script lang="ts">
  import MethodSelect from "./MethodSelect.svelte";
  import HeadersEditor from "./HeadersEditor.svelte";
  import BodyEditor from "./BodyEditor.svelte";
  import { requestStore } from "../../lib/stores/requestStore";
  import { addToHistory } from "../../lib/stores/historyStore";
  import type { GoResponse } from "../../types/request-response";
  // @ts-ignore
  import { Get } from "../../../wailsjs/go/main/App";

  export let isDarkMode: boolean;

  let url = "";
  let method = "GET";
  let headers: { key: string; value: string }[] = [{ key: "", value: "" }];
  let body = "";

  async function sendRequest() {
    const request = { method, url, headers, body };

    try {
      const response: GoResponse = await Get({ method, url, headers, body });
      const responseBody = response.body;

      requestStore.set({ request, body: responseBody });
      addToHistory(request);
      alert(responseBody);
    } catch (error) {
      console.error("Request failed:", error);
    }
  }
</script>

<div
  class="{isDarkMode
    ? 'dark:bg-gray-800 dark:border-gray-600 dark:text-white'
    : 'bg-white border-gray-200 text-black'} p-4 rounded shadow-md mb-4 border"
>
  <div class="flex items-center mb-4">
    <MethodSelect bind:value={method} {isDarkMode} />
    <input
      type="text"
      bind:value={url}
      placeholder="Enter API URL"
      class="flex-1 mx-2 p-2 {isDarkMode
        ? 'dark:bg-gray-700 dark:text-white dark:border-gray-600'
        : 'bg-white text-black border-gray-300'} border rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
    />
    <button
      on:click={sendRequest}
      class="{isDarkMode
        ? 'dark:bg-blue-700 dark:hover:bg-blue-600'
        : 'bg-blue-600 hover:bg-blue-700'} text-white px-4 py-2 rounded"
    >
      Send
    </button>
  </div>
  <HeadersEditor bind:headers {isDarkMode} />
  {#if method !== "GET"}
    <BodyEditor bind:body {isDarkMode} />
  {/if}
</div>

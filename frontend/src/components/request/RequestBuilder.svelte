<script lang="ts">
  import MethodSelect from './MethodSelect.svelte';
  import HeadersEditor from './HeadersEditor.svelte';
  import BodyEditor from './BodyEditor.svelte';
  import { requestStore } from '../../lib/stores/requestStore';
  import { addToHistory } from '../../lib/stores/historyStore';
  import { mockSendRequest } from '../../lib/api/mockData';
  import type { GoResponse } from "../../types/request-response";
  
  // @ts-ignore
  import {Get} from '../../../wailsjs/go/main/App'


  let url = '';
  let method = 'GET';
  let headers = [{ key: '', value: '' }];
  let body = '';

  async function sayHi() {
  const request = { method, url, headers, body };

  const r = {
    body: body,
    headers: headers,
    url: url
  };

  try {
    const response: GoResponse = await Get(r);
    const body = response.body

    alert(body)

    requestStore.set({ request, body });
    addToHistory(request);
  } catch (error) {
    console.error("Request failed:", error);
  }
}


</script>

<div class="bg-white p-4 rounded shadow-md mb-4">
  <div class="flex items-center mb-4">
    <MethodSelect bind:value={method} />
    <input
      type="text"
      bind:value={url}
      placeholder="Enter API URL"
      class="flex-1 mx-2 p-2 border rounded focus:outline-none focus:ring-2 focus:ring-blue-500"
    />
    <button
      on:click={sayHi}
      class="bg-blue-600 text-white px-4 py-2 rounded hover:bg-blue-700"
    >
      Send
    </button>
  </div>
  <HeadersEditor bind:headers />
  {#if method !== "GET"}
    <BodyEditor bind:body />
  {/if}
</div>

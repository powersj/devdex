<script lang="ts">
    export let href = '';
    export let text = '';
    export let icon: any;
</script>

<li>
    <a href={href} class="relative flex flex-row items-center h-11 focus:outline-none hover:bg-gray-50 text-gray-600 hover:text-gray-800 border-l-4 border-transparent hover:border-indigo-500 pr-6">
        <span class="inline-flex justify-center items-center ml-4">
            <svelte:component this={icon} />
        </span>
        <span class="ml-2 tracking-wide truncate">{text}</span>
    </a>
</li>

<!--
Extras for individual items:
    * Blue
        <span class="px-2 py-0.5 ml-auto text-xs font-medium tracking-wide text-indigo-500 bg-indigo-50 rounded-full">New</span>
    * Red
        <span class="px-2 py-0.5 ml-auto text-xs font-medium tracking-wide text-red-500 bg-red-50 rounded-full">New</span>
    * Green
        <span class="px-2 py-0.5 ml-auto text-xs font-medium tracking-wide text-green-500 bg-green-50 rounded-full">New</span>
-->

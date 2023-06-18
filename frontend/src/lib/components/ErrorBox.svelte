<script lang="ts">
    export let text = '';
</script>

{#if text}
    <div class="bg-red-100 border border-red-400 text-red-700 px-4 py-3 rounded relative" role="alert">
        <p>{text}</p>
    </div>
{/if}

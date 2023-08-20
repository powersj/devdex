<script lang="ts">
    export let text = '';

    async function copy() {
        await navigator.clipboard.writeText(text);
    }
</script>


<button class="btn btn-outline" on:click={copy}>
copy
</button>

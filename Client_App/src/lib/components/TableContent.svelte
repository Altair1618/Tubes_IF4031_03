<script lang="ts">
    export let isHeader: Boolean = false
</script>

{#if isHeader}
    <th class="font-normal text-sm p-2"><slot/></th>
{:else}
    <td class="px-2 py-[20px]"><slot/></td>
{/if}
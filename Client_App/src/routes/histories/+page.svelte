<script lang="ts">
	import Button from '$lib/components/ui/button/button.svelte';
    import * as DropdownMenu from '$lib/components/ui/dropdown-menu';
	import { IconDotsCircleHorizontal, IconCircleX, IconCreditCard, IconFileDollar } from '@tabler/icons-svelte';
	import toast from 'svelte-french-toast';
    import TableContent from '$lib/components/TableContent.svelte';
    import { DotsVertical } from 'radix-icons-svelte';

	export let data;
</script>

<div class="flex flex-1 flex-col gap-4 w-full py-4">
	<h2 class="text-2xl">Booking History</h2>
	<div class="grid border border-gray-200 rounded-sm w-fit md:w-full">
        <table class="text-left w-[710px] md:w-full">
            <colgroup>
                <col span="1" class="w-[15%]"/>
                <col span="1" class="w-[20%]"/>
                <col span="1" class="w-[20%]"/>
                <col span="1" class="w-[20%]"/>
                <col span="1" class="w-[10%]"/>
                <col span="1" class="w-[10%]"/>
                <col span="1" class="w-[5%]"/>
            </colgroup>
            <thead>
                <tr class="align-top">
                    <TableContent isHeader>ID</TableContent>
                    <TableContent isHeader>Date</TableContent>
                    <TableContent isHeader>Event Name</TableContent>
                    <TableContent isHeader>Event Time & Location</TableContent>
                    <TableContent isHeader>Amount</TableContent>
                    <TableContent isHeader>Status</TableContent>
                    <TableContent isHeader>Action</TableContent>
                </tr>
            </thead>
            <tbody>
                {#each data.data ?? [] as history, idx}
                    <tr class={`${idx % 2 === 1 ? 'bg-gray-50' : ''} text-sm border-b-[1px] border-black align-top`}>
                        {idx}:
                        <TableContent>{history.go}</TableContent>
                        <TableContent>{history_group.date}</TableContent>
                        <TableContent>{history_group.event_name}</TableContent>
                        <TableContent>{history_group.total_price}</TableContent>
                        <TableContent>{history_group.overall_status}</TableContent>
                        <TableContent>
                            <DropdownMenu.Root preventScroll={false}>
                                <DropdownMenu.Trigger>
                                    <Button class="p-0" variant="ghost">
                                        <DotsVertical size={16} />
                                    </Button>
                                </DropdownMenu.Trigger>
                                <DropdownMenu.Content>
                                    <DropdownMenu.Group>
                                        <DropdownMenu.Item>
                                            <a href={`/histories/group/${history_group.group_id}`} class="flex items-center gap-2">
                                                <IconDotsCircleHorizontal size={18} />
                                                See details
                                            </a>
                                        </DropdownMenu.Item>
                                        {#if true}
                                            <DropdownMenu.Item>
                                                <a href={history_group.payment_url} class="flex items-center gap-2">
                                                    <IconCreditCard size={18} />
                                                    Purchase
                                                </a>
                                            </DropdownMenu.Item>
                                        {:else if true}
                                            <DropdownMenu.Item>
                                                <button class="flex items-center gap-2">
                                                    <IconCircleX size={18} />
                                                    Cancel booking
                                                </button>
                                            </DropdownMenu.Item>
                                        {:else}
                                            <DropdownMenu.Item>
                                                <button class="flex items-center gap-2">
                                                    <IconFileDollar size={18} />
                                                    See invoice
                                                </button>
                                            </DropdownMenu.Item>
                                        {/if}
                                    </DropdownMenu.Group>
                                </DropdownMenu.Content>
                            </DropdownMenu.Root>
                        </TableContent>
                    </tr>
                {/each}
            </tbody>
        </table>
        {#if data.data.length > 0}
            <div class="w-full flex flex-row justify-center py-[20px]">
                <span class="text-seven-font-size-table-content">No booking history</span>
            </div>
        {/if}
    </div>
</div>

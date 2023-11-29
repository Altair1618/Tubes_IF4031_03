<script lang="ts">
	import Button from '$lib/components/ui/button/button.svelte';
    import * as DropdownMenu from '$lib/components/ui/dropdown-menu';
	import { IconDotsCircleHorizontal, IconCircleX, IconCreditCard, IconFileDollar } from '@tabler/icons-svelte';
	import toast from 'svelte-french-toast';
    import TableContent from '$lib/components/TableContent.svelte';
    import { DotsVertical } from 'radix-icons-svelte';
    import { BookingStatus } from '$lib/types/booking.js';
    import { getDateTimeString } from '$lib/utils';
    import { invalidateAll } from '$app/navigation';
	import { enhance as enhance2 } from '$app/forms';
    import type { HistoryResponseData } from '$lib/types/booking';
    import InfiniteScroll from '$lib/components/InfiniteScroll.svelte';

	export let data;

    let histories = data.data

    

</script>

<div class="flex flex-1 flex-col gap-4 w-full py-4">
	<h2 class="text-2xl">Booking History</h2>
	<div class="grid border border-gray-200 rounded-sm w-full overflow-x-auto">
        <table class="text-left min-w-[800px]">
            <colgroup>
                <col span="1" class="w-[10%]"/>
                <col span="1" class="w-[15%]"/>
                <col span="1" class="w-[20%]"/>
                <col span="1" class="w-[20%]"/>
                <col span="1" class="w-[10%]"/>
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
                    <TableContent isHeader>Seat ID</TableContent>
                    <TableContent isHeader>Price</TableContent>
                    <TableContent isHeader>Status</TableContent>
                    <TableContent isHeader>Action</TableContent>
                </tr>
            </thead>
            <tbody>
                {#each histories ?? [] as history, idx}
                    <tr class={`${idx % 2 === 1 ? 'bg-gray-50' : ''} text-sm border-b-[1px] border-gray-200 align-top`}>
                        <TableContent>{history.id}</TableContent>
                        <TableContent>{getDateTimeString(history.createdAt)}</TableContent>
                        <TableContent>{history.eventName}</TableContent>
                        <TableContent>{`${getDateTimeString(history.eventTime)}\n${history.location}`}</TableContent>
                        <TableContent>{history.seatId}</TableContent>
                        <TableContent>{history.price}</TableContent>
                        <TableContent>{history.status}</TableContent>
                        <TableContent>
                            <DropdownMenu.Root preventScroll={false}>
                                <DropdownMenu.Trigger class="w-full flex flex-row justify-center" disabled="{history.status === BookingStatus.PURCHASING || (history.status === BookingStatus.FAILED && !history.report)}">
                                    <Button class="p-0 w-full flex flex-grow justify-center" variant="ghost" disabled="{history.status === BookingStatus.PURCHASING || (history.status === BookingStatus.FAILED && !history.report)}">
                                        <DotsVertical size={16} />
                                    </Button>
                                </DropdownMenu.Trigger>
                                <DropdownMenu.Content>
                                    <DropdownMenu.Group>
                                        {#if history.status === BookingStatus.WAITING_FOR_PAYMENT && history.paymentUrl}
                                            <DropdownMenu.Item>
                                                <a href={history.paymentUrl} class="flex items-center gap-2">
                                                    <IconCreditCard size={18} />
                                                    Purchase
                                                </a>
                                            </DropdownMenu.Item>
                                        {/if}
                                        {#if history.status === BookingStatus.IN_QUEUE || history.status === BookingStatus.WAITING_FOR_PAYMENT}
                                            <DropdownMenu.Item>
                                                <form
                                                    action="?/cancelBooking"
                                                    method="post"
                                                    use:enhance2={() => {
                                                        return async ({ result }) => {
                                                            if (result.type === 'error') {
                                                                toast.error(result.error.message);
                                                            } else if (result.type === 'success') {
                                                                
                                                                histories.forEach(elmt => {
                                                                    if (elmt.id === history.id)
                                                                    {
                                                                        //@ts-ignore
                                                                        elmt.status = result.data?.newStatus
                                                                    }
                                                                })

                                                                histories = histories;
                                                                toast.success('Booking cancelled');
                                                                await invalidateAll();
                                                            }
                                                        };
                                                    }}
                                                >
                                                <input name="id" hidden type="text" value="{history.id}"/>
                                                <button class="flex items-center text-xs gap-2" type="submit">
                                                    <IconCircleX size={18} />
                                                    Cancel booking
                                                </button>
                                                </form>
                                            </DropdownMenu.Item>
                                        {/if}
                                        {#if (history.status === BookingStatus.SUCCESS || history.status === BookingStatus.FAILED) && history.report}
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
        {#if data.data.length === 0}
            <div class="w-full flex flex-row justify-center py-[20px]">
                <span class="text-seven-font-size-table-content">No booking history</span>
            </div>
        {/if}
    </div>
</div>

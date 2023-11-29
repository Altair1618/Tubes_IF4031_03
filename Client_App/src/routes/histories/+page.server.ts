import { PUBLIC_CLIENT_SERVICE_BASE_URL } from '$env/static/public';
import { error, fail } from '@sveltejs/kit';
import type { Actions, PageServerLoad } from './$types';
import type { ClientServiceResponse } from '$lib/types/common';
import { message, setError, superValidate } from 'sveltekit-superforms/server';
import type { CancelBookingResponseData, HistoryResponseData } from '$lib/types/booking';
import { cancelBookingSchema, movePageSchema, purchaseSchema } from '$lib/dto/booking/cancelBooking.dto';

export const load: PageServerLoad = async ({ locals, cookies }) => {
	const response = await fetch(`${PUBLIC_CLIENT_SERVICE_BASE_URL}/bookings`, {
		method: 'GET',
		headers: {
			Authorization: `Bearer ${cookies.get('mikuuuu')}`
		},
		credentials: 'include'
	});
	const responseData: ClientServiceResponse<HistoryResponseData[]> = await response.json();

	if (!response.ok) {
		throw error(response.status, responseData.message);
	}

	return {
		...responseData
	};
};

export const actions = {
	cancelBooking: async ({ request, cookies }) => {
		const form = await superValidate(request, cancelBookingSchema);

		if (!form.valid) {
			return fail(400, { form });
		}

		try {

			const response = await fetch(`${PUBLIC_CLIENT_SERVICE_BASE_URL}/bookings/${form.data.id}/status/cancel`, {
				method: 'PATCH',
				headers: {
					Authorization: `Bearer ${cookies.get('mikuuuu')}`,
					'Content-Type': 'application/json'
				},
				credentials: 'include'
			});

			const responseData: ClientServiceResponse<CancelBookingResponseData> = await response.json();

			if (!response.ok) {
				return fail(responseData.code, { message: responseData.message })
			}

			return {
				message: responseData.message,
				newStatus: responseData.data.newStatus
			};

		} catch (e) {
			console.log(e);
			return fail(500, { message: "Something went wrong, please try again later" })
		}
	},
	movePage: async ({ cookies, request }) => {
		const form = await superValidate(request, movePageSchema);

		if (!form.valid) {
			return fail(400, { form });
		}
		console.log(form.data.page)

		const response = await fetch(`${PUBLIC_CLIENT_SERVICE_BASE_URL}/bookings?page=${form.data.page}`, {
			method: 'GET',
			headers: {
				Authorization: `Bearer ${cookies.get('mikuuuu')}`
			},
			credentials: 'include'
		});
		const responseData: ClientServiceResponse<HistoryResponseData[]> = await response.json();
	
		if (!response.ok) {
			throw error(response.status, responseData.message);
		}

		return {
			message: responseData.message,
			histories: responseData.data
		};
	},

	doPayment: async ({ cookies, request }) => {
		const form = await superValidate(request, purchaseSchema);

		if (!form.valid) {
			return fail(400, { form });
		}

		const response = await fetch(`${PUBLIC_CLIENT_SERVICE_BASE_URL}/payment`, {
			method: 'POST',
			headers: {
				Authorization: `Bearer ${cookies.get('mikuuuu')}`
			},
			body: JSON.stringify({
				payment_url: form.data.payment_url
			}),
			credentials: 'include'
		});
		const responseData: ClientServiceResponse<HistoryResponseData[]> = await response.json();
	
		if (!response.ok) {
			throw error(response.status, responseData.message);
		}

		return {
			message: responseData.message,
			histories: responseData.data
		};
	},
} satisfies Actions;

import { PUBLIC_CLIENT_SERVICE_BASE_URL } from '$env/static/public';
import { error, fail } from '@sveltejs/kit';
import type { Actions, PageServerLoad } from './$types';
import type { ClientServiceResponse } from '$lib/types/common';
import { message, setError, superValidate } from 'sveltekit-superforms/server';
import { updateProfileSchema } from '$lib/dto/profile/updateProfile.dto';

interface HistoryResponseData {
	id: string,
	ticketId: string,
	status: string,
    paymentUrl: string | null,
	report: string | null,
	createdAt: string,
	seatId: string,
	price: number,
	eventName: string,
	eventTime: string,
	location: string
	totalPage: number
}

export const load: PageServerLoad = async ({ locals, cookies }) => {
	return {
		code: 200,
		message: 'yes',
		data: []
	}
	const response = await fetch(`${PUBLIC_CLIENT_SERVICE_BASE_URL}/bookings?user_id=${locals.user?.userId}`, {
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
	
} satisfies Actions;

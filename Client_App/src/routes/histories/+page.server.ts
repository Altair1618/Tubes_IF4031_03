import { PUBLIC_CLIENT_SERVICE_BASE_URL } from '$env/static/public';
import { error, fail } from '@sveltejs/kit';
import type { Actions, PageServerLoad } from './$types';
import type { ClientServiceResponse } from '$lib/types/common';
import { message, setError, superValidate } from 'sveltekit-superforms/server';
import { updateProfileSchema } from '$lib/dto/profile/updateProfile.dto';

interface History {
	groupId: string,
	date: string,
	eventName: string,
	totalPrice: number,
	overallStatus: string,
	paymentUrl: string
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
	const responseData: ClientServiceResponse<History[]> = await response.json();

	if (!response.ok) {
		throw error(response.status, responseData.message);
	}

	return {
		...responseData
	};
};

export const actions = {
	
} satisfies Actions;

import { PUBLIC_CLIENT_SERVICE_BASE_URL } from '$env/static/public';
import { error, fail } from '@sveltejs/kit';
import type { Actions, PageServerLoad } from './$types';
import type { ClientServiceResponse } from '$lib/types/common';
import { message, setError, superValidate } from 'sveltekit-superforms/server';
import { updateProfileSchema } from '$lib/dto/profile/updateProfile.dto';

type Histories = [{
	group_id: string,
	date: string,
	event_name: string,
	total_price: Number,
	overall_status: string,
	payment_url: string
}]

export const load: PageServerLoad = async ({ locals, cookies }) => {
	return {
		code: 200,
		message: 'yes',
		data: []
	}
	const response = await fetch(`${PUBLIC_CLIENT_SERVICE_BASE_URL}/histories`, {
		method: 'GET',
		headers: {
			Authorization: `Bearer ${cookies.get('mikuuuu')}`
		},
		credentials: 'include'
	});
	const responseData: ClientServiceResponse<Histories> = await response.json();

	if (!response.ok) {
		throw error(response.status, responseData.message);
	}

	return {
		...responseData
	};
};

export const actions = {
	
} satisfies Actions;

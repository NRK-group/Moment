import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
import config from '../../../config';

export default function Validation(auth) {
    const navigate = useNavigate();
    useEffect(() => {
        fetch(config.api + '/validate', {
            credentials: 'include',
        }).then(async (resp) => {
            const response = await resp.text();
            if (response !== 'Validated') {
                auth(false);
                navigate('/');
                return;
            }
            auth(true);
            return;
        });
    }, []);
}

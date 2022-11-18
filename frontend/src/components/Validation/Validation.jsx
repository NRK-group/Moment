import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
export default function Validation(auth) {
    const navigate = useNavigate();
    useEffect(() => {
        fetch('http://localhost:5070/validate', {
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

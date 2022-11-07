import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';

export default function ValidRedirect() {
    const navigate = useNavigate();
    //Query the endpoint and check if the cookie present is valid
    useEffect(() => {
        fetch('http://localhost:5070/validate', { credentials: 'include' })
            .then((resp) => {
                return resp.text();
            })
            .then((response) => {
                if (response === 'Validated') {
                    navigate('/home')
                }
            });
    }, []);
}

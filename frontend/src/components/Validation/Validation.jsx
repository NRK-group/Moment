import { useEffect, useState } from 'react';
import { useNavigate } from 'react-router-dom';
export default function Validation(auth) {
    const [authorised, setAuthorised] = useState(false);
    const navigate = useNavigate();
    useEffect(() => {
        fetch('http://localhost:5070/validate', {
            credentials: 'include',
        }).then(async (resp) => {
            const response = await resp.text();
            console.log("RESULT WHILE VALIDATING:   ", response)
            if (response !== 'Validated') {
                setAuthorised(false);
                navigate('/');
                return;
            }
            setAuthorised(true);
            return;
        });
    }, [auth]);
    return authorised;
}

import { useState, useEffect } from 'react';
import Card from '../card/Card'
import './Avatar.css';

function Avatar({ avatarSrc, avatarAlt, styleName, name , avatarRedirect, avatarIcon, cardStyleName }) {
    const [src, setSrc] = useState(avatarSrc);
    const [alt, setAlt] = useState(avatarAlt);
    const [nameS, setNameS] = useState(name);
    const [avatarIconS, setAvatarIconS] = useState(name);
    const [styleNameS, setStyleNameS] = useState(styleName);
    const [cardStyleNameS, setCardStyleNameS] = useState(cardStyleName);


    useEffect(() => {
        setSrc(avatarSrc)
        setAlt(avatarAlt)
        setStyleNameS(styleName)
        setNameS(name)
        setAvatarIconS(avatarIcon)
        setCardStyleNameS(cardStyleName)
    }, [avatarSrc, avatarAlt, styleName, avatarIcon, cardStyleName]);


    return (
        <Card styleName={cardStyleNameS}>
            <img className={styleNameS} src={src} alt={alt} onClick={()=>{avatarRedirect}} />
          { avatarIconS && (<span className='AvatarIcon'><i className="fa-solid fa-circle-plus"></i></span>)  }
            <p>{nameS}</p>
            
        </Card>
    );
}

export default Avatar;

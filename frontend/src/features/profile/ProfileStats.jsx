import { useNavigate } from 'react-router-dom';
import Card from '../../components/card/Card';

export default function ProfileStats(props) {
    const navigate = useNavigate('');

    return (
        <Card styleName={props.styleName}>
            <Card styleName='postStats stats'>
                <p>{props.posts} posts</p>
            </Card>
            <Card styleName='followersStats stats'>
                <p onClick={() => navigate('/profile/followers')}>
                    {props.followers} followers
                </p>
            </Card>
            <Card styleName='followingStats stats'>
                <p onClick={() => navigate('/profile/following')}>{props.following} following</p>
            </Card>
        </Card>
    );
}

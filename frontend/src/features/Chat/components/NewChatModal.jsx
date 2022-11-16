import Input from '../../../components/Input/Input';
import { CloseIcon } from '../../../components/Icons/Icons';
import MiniUserCard from '../../../components/MiniUserCard/MiniUserCard';
import { useNavigate } from 'react-router-dom';
import { GetFollow } from '../hooks/getFollow';
import { useState } from 'react';
export const NewChatModal = () => {
    const navigate = useNavigate();
    let following = GetFollow();
    const [query, setQuery] = useState('');
    let filteredItems = following.filter((item) => {
        return (
            item.name.toLowerCase().includes(query.toLowerCase()) ||
            item.firstName.toLowerCase().includes(query.toLowerCase()) ||
            item.lastName.toLowerCase().includes(query.toLowerCase())
        );
    });
    return (
        <div
            className='newChatModalContainer'
            onClick={() => {
                navigate(`/messages`);
            }}>
            <div
                className='newChatModal'
                onClick={(e) => {
                    e.stopPropagation();
                }}>
                <div className='newChatModalHeader'>
                    <div className='closeIcon'>
                        <CloseIcon
                            action={() => {
                                navigate(`/messages`);
                            }}
                        />
                    </div>
                    <div className='title'>New Message</div>
                    <div></div>
                </div>
                <div className='newChatModalHeaderSearch'>
                    <div>To:</div>
                    <Input
                        styleName='searchFollowing'
                        placeholder={'Search . . .'}
                        onChange={(e) => {
                            setQuery(e.target.value);
                        }}
                    />
                </div>
                <div className='searchResult scrollbar-hidden'>
                    {filteredItems
                        ? filteredItems.map(
                              ({ id, firstName, lastName, name, img }) => (
                                  <MiniUserCard
                                      propsId={id}
                                      name={firstName + ' ' + lastName}
                                      img={img}>
                                      {name}
                                  </MiniUserCard>
                              )
                          )
                        : null}
                </div>
            </div>
        </div>
    );
};

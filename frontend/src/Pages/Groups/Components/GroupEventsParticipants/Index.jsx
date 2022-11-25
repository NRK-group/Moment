import './GroupEventsParticipants.css'
import TabletableParticipants from './Components/Table';



export default function GroupEventsParticipants({data}) {
    return (
        <div id='GroupEventsParticipants' className='GroupListContainer'>
            <div  className='GroupListDiv'>
            <TabletableParticipants data={data}/>
            </div>
        </div>
    );
}

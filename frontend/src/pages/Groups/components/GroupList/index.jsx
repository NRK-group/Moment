import './GroupList.css'
import Table from '../Table';



export default function GroupList({socket, data}) {
    return (
        <div id='GroupListContainer' className='GroupListContainer'>
            <div  className='GroupListDiv'>
            <Table socket={socket} data={data}/>
            </div>
        </div>
    );
}

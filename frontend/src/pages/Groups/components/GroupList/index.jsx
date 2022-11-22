import './GroupList.css'
import Table from '../Table';



export default function GroupList({data}) {
    return (
        <div id='GroupListContainer' className='GroupListContainer'>
            <div  className='GroupListDiv'>
            <Table data={data}/>
            </div>
        </div>
    );
}

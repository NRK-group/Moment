import Card from '../../components/card/Card';
import React, { useState } from 'react';
import DatePicker from 'react-datepicker';

import 'react-datepicker/dist/react-datepicker.css';

export default function AuthDateInput(props) {
    const [startDate, setStartDate] = useState(new Date());
  
    return (
        // <Card styleName={props.styleName}>
            <DatePicker 
      selected={new Date(new Date().setFullYear(new Date().getFullYear() - 20))} 
      onChange={date => setStartDate(date)} 
      dateFormat='dd/MM/yyyy'
      minDate={new Date(new Date().setFullYear(new Date().getFullYear() - 13)).getFullYear()}
      ref={props.dayRef}
    />
        // </Card>
    );
  }


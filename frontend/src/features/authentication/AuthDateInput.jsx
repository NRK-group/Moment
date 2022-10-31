import Card from "../../components/card/Card"

export default function AuthDateInput(props) {
    const DAYS = [], MONTH_OPTIONS = [], MONTHS=['Jan', 'Feb', 'Mar', 'Apr', 'May', 'Jun', 'Jul', 'Aug', 'Sep', 'Oct', 'Nov', 'Dec'], YEARS = []
    for (let i=1; i<=31; i++){
        DAYS.push(<option key={i}>{i}</option>)
    }

    MONTHS.forEach((elem) => {
        MONTH_OPTIONS.push(<option key={elem}>{elem}</option>)
    })
    

    for (let i= +(new Date().getFullYear()); i >= 1904; i--){
        YEARS.push(<option key={i}>{i}</option>)
    }
  return (
    <Card styleName={props.styleName}>
    <select className={props.daySelector} ref={props.dayRef}>
        {DAYS}
    </select>
    <select className={props.monthSelector} ref={props.monthRef}>
        {MONTH_OPTIONS}
    </select>
    <select className={props.yearSelector} ref={props.yearRef}>
        {YEARS}
    </select>
    </Card>
  )
}

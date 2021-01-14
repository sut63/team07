import React, { useState,useEffect} from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
} from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';
//import Autocomplete from '@material-ui/lab/Autocomplete';
//import Typography from '@material-ui/core/Typography';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import Select from '@material-ui/core/Select';
//import InputAdornment from '@material-ui/core/InputAdornment';
//import { EntCarservice } from '../../api/models/EntCarservice';
import { EntUser } from '../../api/models/EntUser';
import { EntUrgent } from '../../api/models/EntUrgent';
import { EntDistance } from '../../api/models/EntDistance';

const useStyles = makeStyles((theme: Theme) =>
createStyles({
  root: {
    display: 'flex',
    flexWrap: 'wrap',
    justifyContent: 'center',
  },
  margin: {
    margin: theme.spacing(3),
  },
  
  withoutLabel: {
    marginTop: theme.spacing(3),
  },
  textField: {
    width: '25ch',
  },
  paper: {
    marginTop: theme.spacing(1),
    marginBottom: theme.spacing(1),
    marginLeft: theme.spacing(1),
  },
  formControl: {
    width: 400,
  },
}),
);

export default function Create() {
 const classes = useStyles();
 const profile = {thisName: 'ระบบบันทึกการใช้รถพยาบาล' };
 const api = new DefaultApi();
 
 const [users, setUsers] = useState<EntUser[]>([]);
 const [urgents, setUrgents] = useState<EntUrgent[]>([]);
 const [distances, setDistances] = useState<EntDistance[]>([]);
 const [userid, setUser] = useState(Number);
 const [urgentid, setUrgent] = useState(Number);
 const [disid, setDistance] = useState(Number);
 const [datetime, setdatetime] = useState(String);
 const [loading, setLoading] = useState(true);
 const [status, setStatus] = useState(false);
 const [alert, setAlert] = useState(true);
 const [customer, setCustomer] = useState(String);
 const [location, setLocation] = useState(String);
 const [information, setInformation] = useState(String);
 
 useEffect(() => {
   const getUrgent = async () => {
     const res = await api.listUrgent({offset :0});
     setLoading(false);
     setUrgents(res);
     console.log(res);
   };
   getUrgent();

   const getDistance = async () => {
     const res = await api.listDistance({offset :0});
     setLoading(false);
     setDistances(res);
     console.log(res);
   };
   getDistance();

   const getUser = async () => {
    const res = await api.listUser();
    setLoading(false);
    setUsers(res);
    console.log(res);
  };
  getUser();

   const checkJobPosition = async () => {
    const jobdata = JSON.parse(String(localStorage.getItem("jobpositiondata")));
    setLoading(false);
    if (jobdata != "เจ้าหน้าที่โอเปอร์เรเตอร์" ) {
      localStorage.setItem("userdata",JSON.stringify(null));
      localStorage.setItem("jobpositiondata",JSON.stringify(null));
      history.pushState("","","./");
      window.location.reload(false);        
    }
    else{
        setUser(Number(localStorage.getItem("userdata")))
    }
  }
checkJobPosition();

}, [loading]);

const handleCustomerChange = (event: any) => {
  setCustomer(event.target.value as string);
};

const handleLocationChange = (event: any) => {
  setLocation(event.target.value as string);
};

const handleInformationChange = (event: any) => {
  setInformation(event.target.value as string);
};

 const handledatetimeChange = (event: any) => {
   setdatetime(event.target.value as string);
 };

 const handleUrgentchange = (event: React.ChangeEvent<{value: unknown}>) => {
   setUrgent(event.target.value as number);
 };

 const handleUserchange = (event: React.ChangeEvent<{value: unknown}>) => {
  setUser(event.target.value as number);
};
  
 const handleDistancechange = (event: React.ChangeEvent<{value: unknown}>) => {
    setDistance(event.target.value as number);
 };

 const CreateCarservice = async ()=>{
   if ((customer != null) && (customer != "") && (location != null) && (location != "") && (information != null) && (information != "") && (datetime != null) && (datetime != "") && (urgentid != null) && (disid != null)) {
   const carservice ={
     userID: userid,
     urgentID: urgentid,
     distanceID: disid,
     customer: customer,
     location: location,
     information: information,
     datetime: datetime + ":00+07:00",
   };
   const res: any = await api.createCarservice({ carservice: carservice });
            setStatus(true);
            if (res.id != '') {
                setAlert(true);
            }
        }
        else {
            setStatus(true);
            setAlert(false);
        }
  const timer = setTimeout(()=>{
     setStatus(false);
     window.location.reload(false);
   },5000);
    };
 
 return (
    <Page theme={pageTheme.library}>
     <Header
       title={`${profile.thisName}`}
       subtitle="บรื๊นๆๆ"
     ></Header>

     <Content>
       <ContentHeader title="เพิ่มบันทึกการใช้รถ">
       
         {status ? (
           <div>
             {alert ? (
               <Alert severity="success">
                 บันทึกสำเร็จ! กด ย้อนกลับ เพื่อดูข้อมูล
               </Alert>
             ) : (
               <Alert severity="warning" style={{ marginTop: 20 }}>
                 บันทึกไม่สำเร็จ!
               </Alert>
             )}
           </div>
         ) : null}
         
         <div className={classes.margin}>
             <Button
               onClick={() => {
                 CreateCarservice();
               }}
               variant="contained"
               color="primary"
             >
               ยืนยันการบันทึก
             </Button>
             <Button
               style={{ marginLeft: 20 }}
               component={RouterLink}
               to="/Carservicemain"
               variant="contained"
             >
               ย้อนกลับ
             </Button>
           </div>
       </ContentHeader>

       <div className={classes.root}>
          <form noValidate autoComplete="off">
          <div className={classes.paper}><strong>ผู้ใช้บริการ</strong></div>
            <TextField className={classes.textField}
            style={{ width: 400 ,marginLeft:20,marginRight:-10}}
      
              id="customer"
              label=""
              variant="standard"
              color="secondary"
              type="string"
              size="medium"
              value={customer}
              onChange={handleCustomerChange}
            />

            <div className={classes.paper}><strong>สถานที่</strong></div>
            <TextField className={classes.textField}
                style={{ width: 400 ,marginLeft:20,marginRight:-10}}
              id="location"
              label=""
              variant="standard"
              color="secondary"
              type="string"
              size="medium"
              value={location}
              onChange={handleLocationChange}
            />  

            <div className={classes.paper}><strong>การใช้บริการ</strong></div>
            <TextField className={classes.textField}
            style={{ width: 400 ,marginLeft:20,marginRight:-10}}
              id="information"
              label=""
              variant="standard"
              color="secondary"
              type="string"
              size="medium"
              value={information}
              onChange={handleInformationChange}
            />

            <div>
          <FormControl
              className={classes.margin}
              variant="outlined"
            >
              <div className={classes.paper}><strong>ขอบเขตระยะทาง</strong></div>
              <InputLabel id="distance-label"></InputLabel>
              <Select
                labelId="distance-label"
                id="ระยะทาง"
                value={disid}
                onChange={handleDistancechange}
                style={{ width: 400 }}
              >
                {distances.map((item: EntDistance) => (
                  <MenuItem value={item.id}>{item.distance}</MenuItem>
                ))}
              </Select>
            </FormControl>
            </div>

            <div>
          <FormControl
              className={classes.margin}
              variant="outlined"
            >
              <div className={classes.paper}><strong>ความเร่งด่วน</strong></div>
              <InputLabel id="urgent-label"></InputLabel>
              <Select
                labelId="urgent-label"
                id="ความเร่งด่วน"
                value={urgentid}
                onChange={handleUrgentchange}
                style={{ width: 400 }}
              >
                {urgents.map((item: EntUrgent) => (
                  <MenuItem value={item.id}>{item.urgent}</MenuItem>
                ))}
              </Select>
            </FormControl>
            </div>

            <FormControl className={classes.margin} >
                <TextField
                 id="Datetime"
                  label="วัน-เวลา"
                  type="datetime-local"
                  value={datetime}
                  onChange={handledatetimeChange}
                  className={classes.textField}
                  InputLabelProps={{
                   shrink: true,
                 }}
                 style={{ width: 250 }}
        
                />
                </FormControl>   
                
                <div>
                   <FormControl
                  className={classes.margin}
                  variant="outlined"
                  >
                  <TextField
                   id="user"
                   label="เจ้าหน้าที่"
                   type="string"
                   size="medium"
                   value={users.filter((filter:EntUser) => filter.id == userid).map((item:EntUser) => `${item.name}`)}
                   style={{ width: 400 }}
                  />
                  </FormControl>
                </div>   
         </form>
       </div>
     </Content>
   </Page>
 );
}

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
import { DefaultApi } from '../../api/apis';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import Select from '@material-ui/core/Select';
import { EntUser } from '../../api/models/EntUser';
import { EntUrgent } from '../../api/models/EntUrgent';
import { EntDistance } from '../../api/models/EntDistance';
import Swal from 'sweetalert2';

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
 //const [errormessege, setErrorMessege] = useState(String);
 const [alerttype, setAlertType] = useState(String);
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
 const [age, setAge] = useState(Number);
 const [location, setLocation] = useState(String);
 const [information, setInformation] = useState(String);
 const [customererror, setCustomerError] = React.useState('');
 const [ageerror, setAgeError] = React.useState('');
 const [locationerror, setLocationError] = React.useState('');
 const [informationerror, setInformationError] = React.useState('');
 
 const Toast = Swal.mixin({
  toast: true,
  position: 'top-end',
  showConfirmButton: false,
  timer: 3000,
  timerProgressBar: true,
  didOpen: toast => {
    toast.addEventListener('mouseenter', Swal.stopTimer);
    toast.addEventListener('mouseleave', Swal.resumeTimer);
  },
});

const setErrorMessege = (icon: any, title: any) => {
  Toast.fire({
    icon: icon,
    title: title,
  });
}
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

const handleCustomerChange = (event: React.ChangeEvent<{ value: any }>) => {
  const { value } = event.target;
    const validateValue = value
    checkpattern('customer', validateValue)
  setCustomer(event.target.value as string);
};

const handleAgeChange = (event: React.ChangeEvent<{ value: any }>) => {
  const { value } = event.target;
    const validateValue = value
    checkpattern('age', validateValue)
  setAge(event.target.value as number);
};

const handleLocationChange = (event: React.ChangeEvent<{ value: any }>) => {
  const { value } = event.target;
    const validateValue = value
    checkpattern('location', validateValue)
  setLocation(event.target.value as string);
};

const handleInformationChange = (event: React.ChangeEvent<{ value: any }>) => {
  const { value } = event.target;
    const validateValue = value
    checkpattern('information', validateValue)
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

 const validateCustomer = (val: string) => {
   return val.match("^[ก-๙a-zA-Z\\s]+$");
 }

 const validateAge = (val: number) => {
  return val <= 999 && val >=1  ? true:false
}

const validateLocation = (val: string) => {
  return val.match("^[ก-๙0-9a-zA-Z\\s]+$");
}

const validateInformation = (val: string) => {
  return val.match("^[ก-๙0-9a-zA-Z\\s]+$");
}

const checkpattern = (id: string, value:string) => {
  console.log(value);
  switch(id) {
    case 'customer' :
      validateCustomer(value) ? setCustomerError('') : setCustomerError('กรุณากรอกชื่อ-นามสกุลให้ถูกต้อง');
      return;
    case 'age':
      validateAge(Number(value)) ? setAgeError('') : setAgeError('กรุณาระบุอายุให้ถูกต้อง');
      return;
      case 'location':
      validateLocation(value) ? setLocationError('') : setLocationError('กรุณากรอกที่อยู่ให้ถูกต้อง');
    return;
      case 'information':
      validateInformation(value) ? setInformationError('') : setInformationError('กรุณากรอกสาเหตุการใช้รถให้ถูกต้องให้ถูกต้อง');
    return;
    default:
      return;
  }
}

 const CreateCarservice = async ()=>{
  if ((datetime != null) && (datetime != "")){
      const apiUrl = 'http://localhost:8080/api/v1/carservices';
      const carservice = {
        userID: userid,
     urgentID: urgentid,
     distanceID: disid,
     customer: customer,
     age : Number(age),
     location: location,
     information: information,
     datetime: datetime + ":00+07:00",
   };
   const requestOptions = {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(carservice),
  };
      fetch(apiUrl, requestOptions)
        .then(response => response.json())
        .then(data => {
          console.log(data);
          if (data.status === true) {
            Toast.fire({
              icon: 'success',
              title: 'บันทึกข้อมูลสำเร็จ',
            });
          }
          else {
            ErrorCaseCheck(data.error.Name);
            setAlertType("error");
          }  
        });
    }
    else {
      Toast.fire({
        icon: 'error',
        title: 'บันทึกข้อมูลไม่สำเร็จ',
      });
      setAlertType("error");
    }
      setStatus(true);
    
  };
  const ErrorCaseCheck = (field: string) => {
    switch(field) {
      case 'customer':
        setErrorMessege("error","ระบุชื่อให้ถูกต้อง");
        return;
      case 'age':
        setErrorMessege("error","ระบุอายุให้ถูกต้อง");
        return;
      case 'location':
        setErrorMessege("error","ระบุสถานที่ให้ถูกต้อง");
        return;
        case 'information':
        setErrorMessege("error","ระบุข้อมูลให้ถูกต้อง");
        return;
      default:
        setErrorMessege("error","บันทึกข้อมูลไม่สำเร็จ");
        return;
    }
  }
 
 return (
    <Page theme={pageTheme.library}>
     <Header
       title={`${profile.thisName}`}
       subtitle="บรื๊นๆๆ"
     ></Header>

     <Content>
       <ContentHeader title="เพิ่มบันทึกการใช้รถ">

         
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
          <div className={classes.paper}><strong>ชื่อผู้ใช้บริการ</strong></div>
            <TextField className={classes.textField}
              style={{ width: 400 ,marginLeft:20,marginRight:-10}}
              error = {customererror ? true : false}
              id="customer"
              label=""
              variant="standard"
              color="secondary"
              type="string"
              size="medium"
              value={customer}
              helperText = {customererror}
              onChange={handleCustomerChange}
            />  

            <div className={classes.paper}><strong>อายุ</strong></div>
            <TextField className={classes.textField}
                style={{ width: 400 ,marginLeft:20,marginRight:-10}}
              error ={ageerror ? true : false}
              id="age"
              label=""
              variant="standard"
              color="secondary"
              type="string"
              size="medium"
              value={age}
              helperText ={ageerror}
              onChange={handleAgeChange}
            />

            <div className={classes.paper}><strong>สถานที่</strong></div>
            <TextField className={classes.textField}
                style={{ width: 400 ,marginLeft:20,marginRight:-10}}
              error = {locationerror ? true : false}
              id="location"
              label=""
              variant="standard"
              color="secondary"
              type="string"
              size="medium"
              value={location}
              helperText = {locationerror}
              onChange={handleLocationChange}
            />  

            <div className={classes.paper}><strong>การใช้บริการ</strong></div>
            <TextField className={classes.textField}
            style={{ width: 400 ,marginLeft:20,marginRight:-10}}
              error = {informationerror ? true : false}
              id="information"
              label=""
              variant="standard"
              color="secondary"
              type="string"
              size="medium"
              value={information}
              helperText = {informationerror}
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

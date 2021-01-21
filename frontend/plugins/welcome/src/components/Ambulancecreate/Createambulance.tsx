import React, { useState, useEffect } from 'react';

import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
  Link,
} from '@backstage/core';
import { makeStyles, Theme, createStyles, formatMs } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';

import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import SaveIcon from '@material-ui/icons/Save';
import Select from '@material-ui/core/Select';
import Typography from '@material-ui/core/Typography';
import { EntCarbrand } from '../../api/models/EntCarbrand';
import { EntInspectionResult } from '../../api/models/EntInspectionResult';
import { EntInsurance } from '../../api/models/EntInsurance';
import { EntUser } from '../../api/models/EntUser';
import { EntAmbulance } from '../../api/models/EntAmbulance';
import { Link as RouterLink } from 'react-router-dom';


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


export default function AmbulanceCreate() {
  const classes = useStyles();
  const profile = { givenName: 'ยินดีต้อนรับสู่ ระบบลงทะเบียนรถ' };
  const api = new DefaultApi();
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [alert2, setAlerts] = useState(true);
  //เก็บข้อมูลที่จะดึงมา
  const [ambulance, setAmbulance] = useState<EntAmbulance[]>([]);
  const [carbrands, setCarbrands] = useState<EntCarbrand[]>([]);
  const [insurances, setInsurances] = useState<EntInsurance[]>([]);
  const [carstatuses, setCarstatuses] = useState<EntInspectionResult[]>([]);
  const [users, setUsers] = useState<EntUser[]>([]);
  const [registrationerror, setRegistrationerror] = React.useState('');
  const [enginepowererror, setEnginepowererror] = React.useState('');
  const [displacementerror, setDisplacementerror] = React.useState('');
  const [errormessege, setErrorMessege] = useState(String);
  const [alerttype, setAlertType] = useState(String);


  const [loading, setLoading] = useState(true);
  const [date, setDate] = useState(String);

  const [carbrandid, setcarbrand] = useState(Number);
  const [insuranceid, setinsurance] = useState(Number);
  const [enginepowers, setenginepower] = useState(Number);
  const [displacements, setdisplacement] = useState(Number);
  const [carstatusid, setcarstatus] = useState(Number);
  const [registration, setregistration] = useState(String);
  const [userid, setUser] = useState(Number);

  useEffect(() => {
    const getAmbulances = async () => {
      const res = await api.listAmbulance({ limit: 120, offset: 0 });
      setLoading(false);
      setAmbulance(res);
    };
    getAmbulances();

    const getCarbrands = async () => {
 
      const br = await api.listCarbrand();
      setLoading(false);
      setCarbrands(br);
    };
    getCarbrands();

 
    const getUsers = async () => {
 
    const us = await api.listUser();
      setLoading(false);
      setUsers(us);
    };
    getUsers();
 
    const getInsurances = async () => {
 
     const sp = await api.listInsurance();
       setLoading(false);
       setInsurances(sp);
     };
     getInsurances();

     const getCarstatuses = async () => {
 
      const ins = await api.listInspectionresult();
        setLoading(false);
        setCarstatuses(ins);
      };
      getCarstatuses();
      const checkJobPosition = async () => {
        const jobdata = JSON.parse(String(localStorage.getItem("jobpositiondata")));
        setLoading(false);
        if (jobdata != "เจ้าหน้าที่รถพยาบาล" ) {
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
  const validateenginepower = (val: number) => {
    return val <= 100 && val >=80 ? true:false
  
  }

  const validatedisplacement = (val: number) => {
    return val <= 4000 && val >=2900 ? true:false
  }

  const validateRegistration = (val: string) => {
    return val.match("^[0-9]{1}[ก-ฮ]{2}[0-9]{4}$|^[ก-ฮ]{2}[0-9]{4}$");
  }
  const checkPattern  = (id: string, value:string) => {
    console.log(value);
    switch(id) {
      case 'registration':
        validateRegistration(value) ? setRegistrationerror('') : setRegistrationerror('เลขทะเบียนขึ้นต้นด้วยตัวเลข1ตัวหรือตัวก-ฮไม่ไเกิน3ตัวตามด้วยเลข4หลัก');
        return;
      case 'enginepower':
        validateenginepower(Number(value)) ? setEnginepowererror('') : setEnginepowererror('กำลังเครื่องยนต์ต้องอยู่ในช่วง 80 - 100 กิโลวัตต์');
      return;
      case 'displacement':
        validatedisplacement(Number(value)) ? setDisplacementerror('') : setDisplacementerror('ความจุตัวถังต้องอยู่ในช่วง 2900 - 4000 ซีซี');
      return;
        default:
          return;
    }
  }
 
  console.log(userid)
  const DatehandleChange = (event: any) => {
    setDate(event.target.value as string);
  };

  const enginehandleChange = (event: React.ChangeEvent<{ value: any }>) => {
    const { value } = event.target;
    const validateValue = value
    checkPattern('enginepower', validateValue)
    setenginepower(event.target.value as number);
    
  };
  const displacementhandleChange = (event: React.ChangeEvent<{ value: any }>) => {
    const { value } = event.target;
    const validateValue = value
    checkPattern('displacement', validateValue)
    setdisplacement(event.target.value as number);
  };
  const CarbrandhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setcarbrand(event.target.value as number);
  };

  const InsurancehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setinsurance(event.target.value as number);
  };

  const statushandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setcarstatus(event.target.value as number);
  };

  const listamb = () => {
    setStatus(false);
    if(errormessege == "บันทึกข้อมูลสำเร็จ"){
    window.location.href ="http://localhost:3000/mainambulance";
    }
  };

 
  const Registrationhandlehange = (event: React.ChangeEvent<{ id?: string; value: any }>) => {
    const id = event.target.id as  typeof registration;
    const { value } = event.target;
    const validateValue = value.toString()
    checkPattern(id, validateValue)
    setregistration(event.target.value as string);
    
  };
  
  
  const forcheck = () => {
    var i = 0;
    var check = 0;
    for (const amb of ambulance){
      console.log(amb);
     i++;
    }
    console.log(i);
    if(i === 0){
      CreateAmbulance();
   }
   else{
    for (const amb of ambulance){
      if(registration === amb.carregistration){
        //console.log("ซ้ำ");
        setAlertType("error");
        setErrorMessege("มีข้อมูลรถในระบบแล้ว")
        setStatus(true);
        check = 1;
        break;
        //window.location.reload(false);
     }
 }
 if(check != 1){
  CreateAmbulance();
}
   }
  }
 
  const checkCaseSaveError = (field: string) => {
    if (field == "carregistration") { setErrorMessege("ข้อมูลfield เลขทะเบียนรถผิด"); }
        else if (field == "enginepower") { setErrorMessege("ข้อมูลfield กำลังเครื่องยนต์ผิด"); }
        else if (field == "displacement") { setErrorMessege("ข้อมูลfield ความจุตัวถังผิด"); }
        else { setErrorMessege("บันทึกไม่สำเร็จใส่ข้อมูลไม่ครบ"); }
  }
  const CreateAmbulance = async ()=>{
  if((date!= null) && (date != "")){
    const apiUrl = 'http://localhost:8080/api/v1/ambulances';
    const ambulance ={
      carbrandID: carbrandid,
      carstatusID: carstatusid,
      insuranceID: insuranceid,
      userID: userid,
      enginepower:Number(enginepowers),
      displacement:Number(displacements),
      registration: registration,
      datetime: date + ":00+07:00"
    };
    const requestOptions = {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify(ambulance),
    };

    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data);
        setStatus(true);
        if (data.status === true) {
          setErrorMessege("บันทึกข้อมูลสำเร็จ");
          setAlertType("success");

      }
      else {
        checkCaseSaveError(data.error.Name);
          setAlertType("error");
      }
  });
}
else{
  setErrorMessege("บันทึกไม่สำเร็จใส่ข้อมูลไม่ครบ");
  setAlertType("error");
  setStatus(true);
}
     };
     
        


  return (
 <Page theme={pageTheme.home}>
      <Header
        title={`${profile.givenName}`}
      >
       <table>
       <tr>
         <th>
           <h3 style={
             {color: "#483D8B",
             background: 'linear-gradient(45deg, #FFFACD 15%, #9932CC 120%)',
             borderRadius: 10,
             height: 20,
             padding: '0 30px',
             
            }
          }>
            {users.filter((filter: EntUser) => filter.id == userid).map((item: EntUser) => `${item.name} (${item.email})`)}
        </h3>
        </th>
        <th>
        &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
          <Link component={RouterLink} to="/maindriver">
            <Button variant="contained"style={{ background: 'linear-gradient(45deg, #FFFACD 15%, #9932CC 120%)',height: 40}}>
              <h3
              style={
                {color: "#483D8B",
                background: 'linear-gradient(45deg, #FFFACD 15%, #9932CC 120%)',
                borderRadius: 10,
                height: 25,
                padding: '0 20px',
                
               }
             }>
               กลับหน้าหลัก
            </h3>
            </Button>
          </Link>
          </th>
       </tr>
       </table>

      </Header>
      <Content>
      <ContentHeader title="เพิ่มข้อมูลรถเข้าสู่ระบบ" >
      {status ? (
                        <div>
                            {alerttype != "" ? (
                                <Alert severity={alerttype} onClose={() => { listamb() }}>
                                    {errormessege}
                                </Alert>
                            ) : null}
                        </div>
                    ) : null}
                </ContentHeader>
        <div className={classes.root}>
          <form noValidate autoComplete="off">
          <div className={classes.paper}><strong>เลขทะเบียนรถ</strong></div>

            <TextField className={classes.textField}
            style={{ width: 400 ,marginLeft:20,marginRight:-10}}
      
              id="registration"
              error = {registrationerror ? true : false}
              label=""
              variant="standard"
              color="secondary"
              type="string"
              size="medium"
              helperText= {registrationerror}
              value={registration}
              onChange={Registrationhandlehange}
            />
            
          <div>
          <FormControl
              className={classes.margin}
              variant="outlined"
            >
              <div className={classes.paper}><strong>รุ่นของรถ</strong></div>
              <InputLabel id="brand-label"></InputLabel>
              <Select
                labelId="carbrand-label"
                id="carbrand"
                value={carbrandid}
                onChange={CarbrandhandleChange}
                style={{ width: 400 }}
              >
                {carbrands.map((item: EntCarbrand) => (
                  <MenuItem value={item.id}>{item.brand}</MenuItem>
                ))}
              </Select>
            </FormControl>
            </div>
            <div>
                            <FormControl
                                className={classes.margin}
                                variant="outlined"
                            >
                               <div className={classes.paper}><strong>กำลังเครื่องยนต์(กิโลวัตต์)</strong></div>
                                <TextField
                                    id="enginepower"
                                    error = {enginepowererror ? true : false}
                                    helperText= {enginepowererror}
                                    type="number"
                                    size="small"
                                    value={enginepowers}
                                    onChange={enginehandleChange}
                                    style={{ width: 100 }}
                                />
                            </FormControl>
                            <FormControl
                                className={classes.margin}
                                variant="outlined"
                            >
                              <div className={classes.paper}><strong>ความจุตัวถัง(ซีซี)</strong></div>
                                <TextField
                                    id="displacement"
                                    error = {displacementerror ? true : false}
                                    helperText= {displacementerror}
                                    type="number"
                                    size="small"
                                    value={displacements}
                                    onChange={displacementhandleChange}
                                    style={{ width: 100 }}
                                />
                            </FormControl>
                            


                            
                        </div>
            
            <div>
            <FormControl
              className={classes.margin}
              variant="outlined"
            >
              <div className={classes.paper}><strong>บริษัทประกันภัยรถ</strong></div>
              <InputLabel id="company-label"></InputLabel>
              <Select
                labelId="company-label"
                id="company"
                value={insuranceid}
                onChange={InsurancehandleChange}
                style={{ width: 400 }}
              >
                {insurances.map((item: EntInsurance) => (
                  <MenuItem value={item.id}>{item.company}</MenuItem>
                ))}
              </Select>
            </FormControl>
            </div>

            <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <div className={classes.paper}><strong>ชื่อเจ้าหน้าที่รถพยาบาล</strong></div>
                <TextField
                                    id="user"
                                    type="string"
                                    size="medium"
                                    value={users.filter((filter:EntUser) => filter.id == userid).map((item:EntUser) => `${item.name} (${item.email})`)}
                                    style={{ width: 400 }}
                                />
              </FormControl>
            </div>

            <div>
            <FormControl
              className={classes.margin}
              variant="outlined"
            >
               <div className={classes.paper}><strong>สถานะของรถ</strong></div>
              <InputLabel id="status-label"></InputLabel>
              <Select
                labelId="status-label"
                id="status"
                value={carstatusid}
                onChange={statushandleChange}
                style={{ width: 400 }}
              >
                  {carstatuses.filter((filter: any) => filter.edges.jobposition.positionName == "เจ้าหน้าที่รถพยาบาล").map((item: EntInspectionResult) => (
                    <MenuItem value={item.id}>{item.resultName}</MenuItem>
                   ))}
              </Select>
            </FormControl>
            </div>

            
            <div className={classes.paper}><strong>วันที่เวลาบันทึกข้อมูล</strong></div>
            <center> 
                <TextField
                    className={classes.formControl}
                    id="datetime"
                    type="datetime-local"
                    value={date}
                    onChange={DatehandleChange}
                    InputLabelProps={{
                     shrink: true,
                     }}
                 />
                <Typography align ="right"></Typography>
                </center> 
            <div className={classes.margin}>
              <center>
              <Button
                onClick={() => {
                  forcheck();
                }}
                variant="outlined"
                color="primary"
                size="large"
                startIcon={<SaveIcon />}
              >
                บันทึกข้อมูล
             </Button>
              <Button
                style={{ marginLeft: 20 }}
                component={RouterLink}
                to="/mainambulance"
                variant="contained"
                size="large"
              >
                กลับ
             </Button>
             </center>
            </div>
          </form>
        </div>
        
        
      </Content>
    </Page>
  );
}


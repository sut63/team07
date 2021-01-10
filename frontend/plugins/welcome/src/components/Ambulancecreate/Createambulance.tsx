import React, { useState, useEffect } from 'react';
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

import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import FormHelperText from '@material-ui/core/FormHelperText';
import Select from '@material-ui/core/Select';
import Typography from '@material-ui/core/Typography';
import { EntCarbrand } from '../../api/models/EntCarbrand';
import { EntInspectionResult } from '../../api/models/EntInspectionResult';
import { EntInsurance } from '../../api/models/EntInsurance';
import { EntUser } from '../../api/models/EntUser';

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
  const profile = { givenName: 'ยินดีต้อนรับสู่ ระบบลงทะเบียนรถ' };
  const api = new DefaultApi();
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  //เก็บข้อมูลที่จะดึงมา
  const [carbrands, setCarbrands] = useState<EntCarbrand[]>([]);
  const [insurances, setInsurances] = useState<EntInsurance[]>([]);
  const [carstatuses, setCarstatuses] = useState<EntInspectionResult[]>([]);
  const [users, setUsers] = useState<EntUser[]>([]);

  const [loading, setLoading] = useState(true);
  const [date, setDate] = useState(String);

  const [carbrandid, setcarbrand] = useState(Number);
  const [insuranceid, setinsurance] = useState(Number);
  const [carstatusid, setcarstatus] = useState(Number);
  const [registration, setregistration] = useState(String);
  const [userid, setUser] = useState(Number);

  useEffect(() => {

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

  console.log(userid)
  const DatehandleChange = (event: any) => {
    setDate(event.target.value as string);
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

  const UserthandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setUser(event.target.value as number);
  };
  const Registrationhandlehange = (event: any) => {
    setregistration(event.target.value as string);
  };

  const CreateAmbulance = async ()=>{
    if ((registration != null) && (registration != "") && (date!= null) && (date != "") && (carbrandid != null) && (carstatusid != null) && (insuranceid != null) ) {
    const ambulance ={
      carbrandID: carbrandid,
      carstatusID: carstatusid,
      insuranceID: insuranceid,
      userID: userid,
      registration: registration,
      datetime: date + ":00+07:00"
    };
    const res: any = await api.createAmbulance({ ambulance: ambulance });
             setStatus(true);
             if (res.id != '') {
                 setAlert(true);
                 window.location.reload(false);
             }
         }
         else {
             setStatus(true);
             setAlert(false);
         }
         const timer = setTimeout(() => {
             setStatus(false);
         }, 1000);
     };

  return (
 <Page theme={pageTheme.home}>
      <Header
        title={`${profile.givenName}`}
      ></Header>
      <Content>
      <ContentHeader title="เพิ่มข้อมูลรถเข้าสู่ระบบ">
      {status ? (
           <div>
             {alert ? (
               <Alert severity="success">
                 เพิ่มบันทึกเรียบร้อย!
               </Alert>
             ) : (
               <Alert severity="warning" style={{ marginTop: 20 }}>
                 บันทึกไม่สำเร็จ!
               </Alert>
             )}
           </div>
         ) : null}
                </ContentHeader>
        <div className={classes.root}>
          <form noValidate autoComplete="off">
          <div className={classes.paper}><strong>เลขทะเบียนรถ</strong></div>

            <TextField className={classes.textField}
            style={{ width: 400 ,marginLeft:20,marginRight:-10}}
      
              id="registration"
              label=""
              variant="standard"
              color="secondary"
              type="string"
              size="medium"
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

            <div></div>
            
            
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

            <div className={classes.margin}>
              <Button
                onClick={() => {
                  CreateAmbulance();
                }}
                variant="contained"
                color="primary"
              >
                ยืนยัน
             </Button>
              <Button
                style={{ marginLeft: 20 }}
                component={RouterLink}
                to="/mainambulance"
                variant="contained"
              >
                กลับ
             </Button>
            </div>
          </form>
        </div>
        
        
      </Content>
    </Page>
  );
}


import React, { useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
  Content,
  Header,
  Page,
  Link,
  pageTheme,
  ContentHeader,
} from '@backstage/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import TextField from '@material-ui/core/TextField';
import Button from '@material-ui/core/Button';
import FormControl from '@material-ui/core/FormControl';
import { Alert } from '@material-ui/lab';
import { DefaultApi } from '../../api/apis';
import ComponentsTable from '../TransportTable';
import InputLabel from '@material-ui/core/InputLabel';
import MenuItem from '@material-ui/core/MenuItem';
import Select from '@material-ui/core/Select';
import Typography from '@material-ui/core/Typography';
import { EntAmbulance } from '../../api/models/EntAmbulance';
import { EntHospital } from '../../api/models/EntHospital';
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
  const profile = { givenName: 'ยินดีต้อนรับสู่ ระบบขนส่งผู้ป่วย' };
  const api = new DefaultApi();
  const [status, setStatus] = useState(false);
  const [errormessege, setErrorMessege] = useState(String);
  const [alerttype, setAlertType] = useState(String);
  //เก็บข้อมูลที่จะดึงมา
  const [sends, setSends] = useState<EntHospital[]>([]);
  const [receives, setReceives] = useState<EntHospital[]>([]);
  const [ambulances, setAmbulances] = useState<EntAmbulance[]>([]);
  const [users, setUsers] = useState<EntUser[]>([]);
  const [symptomerror, setSymptomerror] = React.useState('');
  const [drugallergyerror, setdrugallergyerror] = React.useState('');
  const [noteerror, setnoteerror] = React.useState('');
  const [loading, setLoading] = useState(true);

  const [drugallergy, setDrugallergy] = useState(String);
  const [symptom, setSymptom] = useState(String);
  const [note, setNote] = useState(String);
  const [sendid, setsend] = useState(Number);
  const [receiveid, setreceive] = useState(Number);
  const [ambulanceid, setAmbulance] = useState(Number);
  const [userid, setUser] = useState(Number);

  useEffect(() => {

    const getSends = async () => {

      const sd = await api.listHospital();
      setLoading(false);
      setSends(sd);
    };
    getSends();

    const getUsers = async () => {
      const us = await api.listUser();
      setLoading(false);
      setUsers(us);
    };
    getUsers();

    const getReceives = async () => {
      const rs = await api.listHospital();
      setLoading(false);
      setReceives(rs);
    };
    getReceives();


    const getAmbulances = async () => {
      const am = await api.listAmbulance();
      setLoading(false);
      setAmbulances(am);
    };
    getAmbulances();

    const checkJobPosition = async () => {
      const jobdata = JSON.parse(String(localStorage.getItem("jobpositiondata")));
      setLoading(false);
      if (jobdata != "เจ้าหน้าที่รถพยาบาล") {
        localStorage.setItem("userdata", JSON.stringify(null));
        localStorage.setItem("jobpositiondata", JSON.stringify(null));
        history.pushState("", "", "./");
        // window.location.reload(false);        
      }
      else {
        setUser(Number(localStorage.getItem("userdata")))
      }
    }
    checkJobPosition();

  }, [loading]);

  console.log(userid)


  const SendhandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setsend(event.target.value as number);
  };

  const ReceivehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setreceive(event.target.value as number);
  };

  const AmbulancehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
    setAmbulance(event.target.value as number);
  };
  const SymptomhandleChange = (event: React.ChangeEvent<{ value: any }>) => {
    const { value } = event.target;
    const validateValue = value
    checkpattern('symptom', validateValue)
    setSymptom(event.target.value as string);
  };
  const DrugallergyhandleChange = (event: React.ChangeEvent<{ value: any }>) => {
    const { value } = event.target;
    const validateValue = value
    checkpattern('drugallergy', validateValue)
    setDrugallergy(event.target.value as string);
  };
  const NotehandleChange = (event: React.ChangeEvent<{ value: any }>) => {
    const { value } = event.target;
    const validateValue = value
    checkpattern('note', validateValue)
    setNote(event.target.value as string);
  };
  const validateSymptom = (val: string) => {
    return val.match("^[ก-๙0-9a-zA-Z\\s-]+$");
  }
  const validateDrugallergy = (val: string) => {
    return val.match("^[ก-๙0-9a-zA-Z\\s-]+$");
  }
  const validateNote = (val: string) => {
    return val.match("^[ก-๙0-9a-zA-Z\\s-]+$");
  }
  const checkpattern = (id: string, value:string) => {
    console.log(value);
    switch(id) {
      case 'symptom':
        validateSymptom(value) ? setSymptomerror('') : setSymptomerror('กรุณณาใส่ข้อมูลให้ถูกต้อง ถ้าไม่มีให้ใส่ - ');
      return;
      case 'drugallergy':
        validateDrugallergy(value) ? setdrugallergyerror('') : setdrugallergyerror('กรุณณาใส่ข้อมูลให้ถูกต้อง ถ้าไม่มีให้ใส่ - ');
      return;
      case 'note':
        //value = nil ? setnoteerror('') : setnoteerror('กรุณณาใส่หมายเหตุเพิ่มเติม ถ้าไม่มีให้ใส่ - ');
        validateNote(value) ? setnoteerror('') : setnoteerror('กรุณณาใส่หมายเหตุเพิ่มเติม ถ้าไม่มีให้ใส่ - ');
      return;
        default: 
          return;
    }
  }
  const CreateTransport = async () => {
    if (sendid != receiveid ) {
      const transports = {
        sendID: sendid,
        receiveID: receiveid,
        ambulanceID: ambulanceid,
        userID: userid,
        drugallergy: drugallergy,
        symptom: symptom,
        note: note,
      };

      console.log(transports);
      const apiUrl = 'http://localhost:8080/api/v1/transports';
      const requestOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(transports),
      };
      fetch(apiUrl, requestOptions)
        .then(response => response.json())
        .then(data => {
          console.log(data);
          if (data.status === true) {
            setErrorMessege("บันทึกข้อมูลสำเร็จ");
            setAlertType("success");

          }
          else {
            ErrorCaseCheck(data.error.Name);
            setAlertType("error");
          }
          
        });
    }
    else {
      ErrorCaseCheck("hospital");
      setAlertType("error");
    }
    setStatus(true);
  };
  const ErrorCaseCheck = (casename: string) => {
    if (casename == "symptom") { setErrorMessege("กรุณาใส่อาการของผู้ป่วย"); }
    else if (casename == "drugallergy") { setErrorMessege("กรุณาใส่การแพ้ยาของผู้ป่วย"); }
    else if (casename == "hospital") { setErrorMessege("ข้อมูลโรงพยาบาลซ้ำกัน"); }
    else if (casename == "note") { setErrorMessege("กรุณณาใส่หมายเหตุเพิ่มเติม"); setnoteerror('กรุณณาใส่หมายเหตุเพิ่มเติม ถ้าไม่มีให้ใส่ - ');}
    else { setErrorMessege("บันทึกไม่สำเร็จ"); }
  }

  return (
    <Page theme={pageTheme.home}>
      <Header
        title={`${profile.givenName}`}
      >
      </Header>
      <Content>
        <ContentHeader title="กรอกข้อมูลส่งตัวผู้ป่วย">
          {status ? (
            <div>
              {alerttype != "" ? (
                <Alert severity={alerttype} onClose={() => { setStatus(false) }}>
                  {errormessege}
                </Alert>
              ) : null}
            </div>
          ) : null}
          <Link component={RouterLink} to="/SearchTransport">
            <Button variant="contained" color="primary" style={{ backgroundColor: "#21b6ae" }}>
              ค้นหาการส่งตัวผู้ป่วย
            </Button>
          </Link>
        </ContentHeader>

        <div className={classes.root}>
          <form noValidate autoComplete="off">

            

            <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <div className={classes.paper}><strong>โรงพยาบาลต้นทาง</strong></div>
                <InputLabel id="send-label"></InputLabel>
                <Select
                  labelId="send-label"
                  id="send"
                  value={sendid}
                  onChange={SendhandleChange}
                  style={{ width: 400 }}
                >
                  {sends.map((item: EntHospital) => (
                    <MenuItem value={item.id}>{item.hospital}</MenuItem>
                  ))}
                </Select>
              </FormControl>
            </div>
            <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <div className={classes.paper}><strong>โรงพยาบาลปลายทาง</strong></div>
                <InputLabel id="receive-label"></InputLabel>
                <Select
                  labelId="receive-label"
                  id="receive"
                  value={receiveid}
                  onChange={ReceivehandleChange}
                  style={{ width: 400 }}
                >
                  {receives.map((item: EntHospital) => (
                    <MenuItem value={item.id}>{item.hospital}</MenuItem>
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
                  value={users.filter((filter: EntUser) => filter.id == userid).map((item: EntUser) => `${item.name} (${item.email})`)}
                  style={{ width: 400 }}
                />
              </FormControl>
            </div>

            <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <div className={classes.paper}><strong>รถพยาบาล</strong></div>
                <InputLabel id="ambulance-label"></InputLabel>
                <Select
                  labelId="ambulance-label"
                  id="ambulance"
                  value={ambulanceid}
                  onChange={AmbulancehandleChange}
                  style={{ width: 400 }}
                >
                  {ambulances.filter((filter: any) => filter.edges.Hasstatus.resultName == "พร้อมใช้งาน").map((item: any) => (
                    <MenuItem value={item.id}>{item.carregistration}</MenuItem>
                  ))}
                </Select>
              </FormControl>
            </div>
            <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <div className={classes.paper}><strong>อาการ</strong></div>
                <TextField
                  id="symptom"
                  error = {symptomerror ? true : false}
                  type="string"
                  size="medium"
                  onChange={SymptomhandleChange}
                  value={symptom}
                  helperText= {symptomerror}
                  style={{ width: 400 }}
                />
              </FormControl>
            </div> <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <div className={classes.paper}><strong>การแพ้ยา</strong></div>
                <TextField
                  id="drugallergy"
                  error = {drugallergyerror ? true : false}
                  type="string"
                  size="medium"
                  onChange={DrugallergyhandleChange}
                  value={drugallergy}
                  helperText= {drugallergyerror}
                  style={{ width: 400 }}
                />
              </FormControl>
            </div> <div>
              <FormControl
                className={classes.margin}
                variant="outlined"
              >
                <div className={classes.paper}><strong>หมายเหตุเพิ่มเติม</strong></div>
                <TextField
                  id="note"
                  error = {noteerror ? true : false}
                  type="string"
                  size="medium"
                  helperText={noteerror}
                  onChange={NotehandleChange}
                  value={note}
                  style={{ width: 400 }}
                />
              </FormControl>
            </div>



            <Typography align="right"></Typography>

            <div className={classes.margin}>
              <Button
                onClick={() => {
                  CreateTransport();
                }}
                variant="contained"
                color="primary"
              >
                ยืนยัน
             </Button>
              <Button
                style={{ marginLeft: 20 }}
                component={RouterLink}
                to="/maindriver"
                variant="contained"
              >
                กลับ
             </Button>
            </div>
          </form>
        </div>
        <ComponentsTable></ComponentsTable>
      </Content>
    </Page>
  );
}

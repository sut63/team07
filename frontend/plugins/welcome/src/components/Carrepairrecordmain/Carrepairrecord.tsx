import React, { useState, useEffect } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
  ApiProvider,
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
import { EntCarInspection } from '../../api/models/EntCarInspection';
import { EntUser } from '../../api/models/EntUser';
import { EntRepairing } from '../../api/models/EntRepairing';
import ComponentsTable from '../CarrepairrecordTable';
import { colors } from '@material-ui/core';
import { red } from '@material-ui/core/colors';

import SearchIcon from '@material-ui/icons/Search';
import SaveIcon from '@material-ui/icons/Save';
import CachedIcon from '@material-ui/icons/Cached';

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

export default function CarInspectionPage() {
    const classes = useStyles();
    const profile = { givenName: 'ระบบบันทึกการบำรุงรักษา' };
    const api = new DefaultApi();

    const [status, setStatus] = useState(false);
    const [errormessage, setErrormessage] = useState(String);
    const [alerttype, setAlertType] = useState(String);
    const [alert, setAlert] = useState(true);
    const [partrepairerror, setPartrepairerror] = React.useState('');
    const [priceerror, setPriceerror] = React.useState('');
    const [techniciancommenterror, setTechniciancommenterror] = React.useState('');

    const [users, setUsers] = useState<EntUser[]>([]);
    const [carinspections, setCarinspections] = useState<EntCarInspection[]>([]);
    const [repairings, setRepairings] = useState<EntRepairing[]>([]);

    const [loading, setLoading] = useState(true);
    const [datetime, setDatetime] = useState(String);
    const [partrepairdata, setPartrepair] = useState(String);
    const [pricedata, setPrice] = useState(Number);
    const [techniciancommentdata, setTechniciancomment] = useState(String);

    const [repairingid, setRepairing] = useState(Number);
    const [carinspectionid, setCarinspection] = useState(Number);
    const [userid, setUser] = useState(Number);

    useEffect(() => {
        const getCarinspections = async () => {
            const ci = await api.listCarinspection({ limit: 10, offset: 0 });
            setLoading(false);
            setCarinspections(ci);
        };
        getCarinspections();

        const getRepairings = async () => {
            const r = await api.listRepairing({limit:10, offset: 0});
            setLoading(false);
            setRepairings(r);
        };
        getRepairings();

        const getUsers = async () => {
            const us = await api.listUser();
            setLoading(false);
            setUsers(us);
        };
        getUsers();
        
        const checkJobPosition = async () => {
            const jobdata = JSON.parse(String(localStorage.getItem("jobpositiondata")));
            setLoading(false);
            if (jobdata != "เจ้าหน้าที่ซ่อมบำรุงรถ" ) {
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

    const validatePartrepair = (val: string) => {
            return val.match("^[ก-๙a-zA-Z-\\s]+$");
    }
    
    const validatePrice = (val: number) => {
            return val > 0 ? true:false 
    }
    
    const validateTechniciancomment = (val: string) => {
            return val.match("^[ก-๙0-9a-zA-Z-\\s]+$")||val.length==0;
    }

    console.log(userid)
    const CarinspectionhandleChange = (event: React.ChangeEvent<{ value: unknown}>) => {
        setCarinspection(event.target.value as number)
    };

    const RepairinghandleChange = (event: React.ChangeEvent<{ value: unknown}>) => {
        setRepairing(event.target.value as number)
    };

    const UserhandleChange = (event: React.ChangeEvent<{ value: unknown}>) => {
        setUser(event.target.value as number)
    };

    const DateTimehandleChange = (event: any) => {
        setDatetime(event.target.value as string);
    };

    const PartrepairhandleChange = (event: React.ChangeEvent<{ value: any }>) => {
        const { value } = event.target;
        const validateValue = value
        checkPattern('partrepair', validateValue)
        setPartrepair(event.target.value as string);
      };

    const PricehandleChange = (event: React.ChangeEvent<{ value: any }>) => {
        const { value } = event.target;
        const validateValue = value
        checkPattern('price', validateValue)
        setPrice(event.target.value as number);
      };

    const TechniciancommenthandleChange = (event: React.ChangeEvent<{ value: any }>) => {
        const { value } = event.target;
        const validateValue = value
        checkPattern('techniciancomment', validateValue)
        setTechniciancomment(event.target.value as string);
      };

    const checkPattern  = (id: string, value:string) => {
        console.log(value);
        switch(id) {
          case 'partrepair':
            validatePartrepair(value) ? setPartrepairerror('') : setPartrepairerror('หมายเหตุไม่ควรใส่ตัวอักษรพิเศษและตัวเลขถ้าไม่มีให้ใส่ -');
            return;
          case 'price':
            validatePrice(Number(value)) ? setPriceerror('') : setPriceerror('กรุณากรอกราคาของการซ่อมให้ถูกต้อง');
          return;
          case 'techniciancomment':
            validateTechniciancomment(value) ? setTechniciancommenterror('') : setTechniciancommenterror('คอมเมนท์ไม่ควรใส่ตัวอักษรพิเศษถ้าไม่มีให้ใส่ -');
          return;
            default:
              return;
        }
      }

    const CreateCarrepairrecord = async () => {
        const carrepairrecord = {
            carInspectionID : carinspectionid,
            repairingID : repairingid,
            userID : userid,
            datetime : datetime + ":00+07:00",
            partrepair : partrepairdata,
            price : Number(pricedata),
            techniciancomment : techniciancommentdata,
        };
        console.log(carrepairrecord);
        const apiUrl = 'http://localhost:8080/api/v1/carrepairrecords';
        const requestOptions = {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(carrepairrecord),
        };

        console.log(carrepairrecord);

        fetch(apiUrl, requestOptions)
            .then(response => response.json())
            .then(data => {
                console.log(data);
                if (data.status === true) {
                    setErrormessage("บันทึกข้อมูลสำเร็จ");
                    setAlertType("success");

                }
                else {
                    ErrorCaseCheck(data.error.Name);
                    setAlertType("error");
                }
            });
        setStatus(true);
        console.log(carrepairrecord);
    };

    const ErrorCaseCheck = (casename: string) => {
        if (casename == "partrepair") { setErrormessage("กรุณากรอกหมายเหตุที่ซ่อมด้วยถ้าไม่มีให้ใส่ -");setPartrepairerror('หมายเหตุไม่ควรใส่ตัวอักษรพิเศษและตัวเลขถ้าไม่มีให้ใส่ -'); }
        else if (casename == "price") { setErrormessage("กรุณากรอกราคาของการซ่อมให้ถูกต้อง"); }
        else if (casename == "techniciancomment") { setErrormessage("คอมเมนท์ไม่ควรใส่ตัวอักษรพิเศษถ้าไม่มีให้ใส่ -"); }
        else { setErrormessage("บันทึกไม่สำเร็จ"); }
      }

    return (
        <Page theme={pageTheme.home}>
            <Header title={`${profile.givenName}`}>
            <div className={classes.margin}>
                    <Button
                        onClick={() => {
                            history.pushState("", "", "/carrepairrecordsearch");
                            window.location.reload(false);
                        }}
                        variant="contained"
                        color="primary"
                        style={{ width: 300, backgroundColor: "#34eb77", color: "#000000", fontSize: 18 }}
                        startIcon={<SearchIcon />}
                    >
                        <b>ค้นหาบันทึกซ่อมบำรุง</b>
                    </Button>
                </div>
            </Header>
            <Content>
                <ContentHeader title="ลงข้อมูลซ่อมบำรุง">
                    {status ? (
                        <div>
                            {alerttype != "" ? (
                                <Alert severity={alerttype} onClose={() => { setStatus(false) }}>
                                    {errormessage}
                                </Alert>
                            ) : null}
                        </div>
                    ) : null}
                </ContentHeader>
                <div className = {classes.root}>
                    <form noValidate autoComplete="off">
                        
                        <div>
                            <FormControl
                                className={classes.margin}
                                variant="outlined"
                            >
                                <div className={classes.paper}><strong>รถพยาบาล</strong></div>
                                <InputLabel id="brand-lavel"></InputLabel>
                                <Select
                                    labelId="carinspection-label"
                                    id="carinspection"
                                    value={carinspectionid}
                                    onChange={CarinspectionhandleChange}
                                    style={{ width: 400}}
                                >
                                    {carinspections.filter((filter:any) => filter.edges?.inspectionresult?.resultName == "ส่งซ่อมแซม").map((item: any) => (
                                        <MenuItem value={item.id}>{item.edges?.ambulance?.carregistration}</MenuItem>
                                    ))}
                                </Select>
                            </FormControl>
                        </div>

                        <div>
                            <FormControl
                                className={classes.margin}
                                variant="outlined"
                            >
                                <div className={classes.paper}><strong>ส่วนที่ซ่อม</strong></div>
                                <InputLabel id="repairing-label"></InputLabel>
                                <Select
                                    labelId="repairing-label"
                                    id="ส่วนที่ซ่อม"
                                    value={repairingid}
                                    onChange={RepairinghandleChange}
                                    style={{ width: 400}}
                                >
                                    {repairings.map((item: EntRepairing) => (
                                        <MenuItem value={item.id}>{item.repairpart}</MenuItem>
                                    ))}
                                </Select>
                            </FormControl>
                        </div>

                        <div>
                            <FormControl
                                className={classes.margin}
                                variant="outlined"
                            >
                                <TextField
                                    id="pathrepair"
                                    error={partrepairerror ? true:false}
                                    label="หมายเหตุส่วนที่ซ่อม"
                                    type="string"
                                    size="medium"
                                    value={partrepairdata}
                                    helperText={partrepairerror}
                                    onChange={PartrepairhandleChange}
                                    style={{ width: 400 }}></TextField>
                            </FormControl>
                            
                        </div>

                        <div>
                            <FormControl>
                            <TextField
                                    id="price"
                                    error={priceerror ? true:false}
                                    label="เงินที่ใช้สำหรับการซ่อม"
                                    type="string"
                                    size="medium"
                                    value={pricedata}
                                    helperText={priceerror}
                                    onChange={PricehandleChange}
                                    style={{ width: 200 }}></TextField>
                            </FormControl>
                            <FormControl
                                className={classes.margin}
                                variant="outlined"
                            >
                                <TextField
                                    id="deathtime"
                                    label="วัน/เดือน/ปี เวลา"
                                    type="datetime-local"
                                    value={datetime}
                                    onChange={DateTimehandleChange}
                                    className={classes.textField}
                                    InputLabelProps={{
                                        shrink: true,
                                    }}
                                    style={{ width: 200 }}
                                />
                            </FormControl>
                        </div>
                        
                        <div>
                            <FormControl
                                className={classes.margin}
                                variant="outlined"
                            >
                                <TextField
                                    id="technicianerror"
                                    error={techniciancommenterror ? true:false}
                                    label="ความคิดเห็นเพิ่มเติมจากช่าง"
                                    type="string"
                                    size="medium"
                                    value={techniciancommentdata}
                                    helperText={techniciancommenterror}
                                    onChange={TechniciancommenthandleChange}
                                    style={{ width: 400 }}></TextField>
                            </FormControl>
                            
                        </div>

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
                                    value={users.filter((filter:EntUser) => filter.id == userid).map((item:EntUser) => `${item.name} (${item.email})`)}
                                    style={{ width: 400 }}
                                />
                            </FormControl>
                        </div>

                        <center>
                        <div className={classes.margin}>
                            <Button
                                onClick={() => {
                                    CreateCarrepairrecord();
                                }}
                                variant="contained"
                                color="primary"
                            >
                                ยืนยัน
                            </Button>
                        </div>
                        </center>
                    </form>
                </div>
                <ComponentsTable></ComponentsTable>
            </Content>
        </Page>   
    );
}
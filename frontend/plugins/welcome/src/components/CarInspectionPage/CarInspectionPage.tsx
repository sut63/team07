import React, { useState, useEffect } from 'react';
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
import Select from '@material-ui/core/Select';

import { EntUser } from '../../api/models/EntUser';
import { EntInspectionResult } from '../../api/models/EntInspectionResult';
import { EntAmbulance } from '../../api/models/EntAmbulance';

import CardMedia from '@material-ui/core/CardMedia';
import { Image1Base64Function } from '../../image/Image1';
import { Image2Base64Function } from '../../image/Image2';

import ComponanceTable from '../CarInspectionTable';

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
        media: {
            height: 0,
            marginLeft: 25,
            maxWidth: 300,
        }

    }),
);


export default function CarInspectionPage() {
    const classes = useStyles();
    const profile = { givenName: 'ยินดีต้อนรับสู่ ระบบตรวจสภาพรถ' };
    const api = new DefaultApi();
    const [status, setStatus] = useState(false);
    const [errormessege, setErrorMessege] = useState(String);
    const [alerttype, setAlertType] = useState(String);

    const [errorWheelCenter, setErrorWheelCenter] = useState(true);
    const [errorBlackSmoke, setErrorBlackSmoke] = useState(true);
    const [errorSoundLevel, setErrorSoundLevel] = useState(true);

    const [users, setUsers] = useState<EntUser[]>([]);
    const [ambulances, setAmbulances] = useState<EntAmbulance[]>([]);
    const [inspectionresults, setInspectionResults] = useState<EntInspectionResult[]>([]);

    const [loading, setLoading] = useState(true);
    const [datetime, setDatetime] = useState(String);
    const [note, setNote] = useState(String);

    const [wheelcenterdata, setWheelCenter] = useState(Number);
    const [blacksmokedata, setBlackSmoke] = useState(Number);
    const [soundleveldata, setSoundLevel] = useState(Number);

    const [inspectionresultid, setInspectionResult] = useState(Number);
    const [ambulanceid, setAmbulance] = useState(Number);
    const [userid, setUser] = useState(Number);

    useEffect(() => {
        const getAmbulances = async () => {
            const res = await api.listAmbulance({ limit: 10, offset: 0 });
            setLoading(false);
            setAmbulances(res);
        };
        getAmbulances();

        const getInspectionResults = async () => {
            const res = await api.listInspectionresult();
            setLoading(false);
            setInspectionResults(res);
        };
        getInspectionResults();

        const getUsers = async () => {
            const res = await api.listUser();
            setLoading(false);
            setUsers(res);
        };
        getUsers();

        const checkJobPosition = async () => {
            const jobdata = JSON.parse(String(localStorage.getItem("jobpositiondata")));
            setLoading(false);
            if (jobdata != "เจ้าหน้าที่ตรวจสภาพรถ") {
                localStorage.setItem("userdata", JSON.stringify(null));
                localStorage.setItem("jobpositiondata", JSON.stringify(null));
                history.pushState("", "", "./");
                window.location.reload(false);
            }
            else {
                setUser(Number(localStorage.getItem("userdata")))
            }
        }
        checkJobPosition();

    }, [loading]);

    const DateTimehandleChange = (event: any) => {
        setDatetime(event.target.value as string);
    };

    const SoundLevelhandleChange = (event: any) => {
        setSoundLevel(event.target.value as number);
        ValidateSoundLevel(event.target.value as number)
    };

    const BlackSmokehandleChange = (event: any) => {
        setBlackSmoke(event.target.value as number);
        ValidateBlackSmoke(event.target.value as number)
    };

    const WheelCenterhandleChange = (event: any) => {
        setWheelCenter(event.target.value as number);
        ValidateWheelCenter(event.target.value as number)
    };

    const NotehandleChange = (event: any) => {
        setNote(event.target.value as string);
    };

    const InspectionResulthandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setInspectionResult(event.target.value as number);
    };

    const AmbulancehandleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
        setAmbulance(event.target.value as number);
    };

    const CreateCarInspection = async () => {
        if (datetime != "") {
            const carinspection = {
                blackSmoke: Number(blacksmokedata),
                soundLevel: Number(soundleveldata),
                wheelCenter: Number(wheelcenterdata),
                inspectionResultID: inspectionresultid,
                ambulanceID: ambulanceid,
                userID: userid,
                datetime: datetime + ":00+07:00",
                note: note


            };
            console.log(carinspection);
            const apiUrl = 'http://localhost:8080/api/v1/carinspections';
            const requestOptions = {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(carinspection),
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
            ErrorCaseCheck("เวลาไม่ได้ใส่");
            setAlertType("error");
        }

        setStatus(true);
    };

    const ErrorCaseCheck = (casename: string) => {
        ValidateWheelCenter(Number(wheelcenterdata));
        ValidateSoundLevel(Number(soundleveldata));
        ValidateBlackSmoke(Number(blacksmokedata));
        if (casename == "wheel_center") { setErrorMessege("ค่า ศูนย์ล้อ ต้องมากกว่า 0"); }
        else if (casename == "sound_level") { setErrorMessege("ค่า ระดับเสียง ต้องมากกว่า 0"); }
        else if (casename == "blacksmoke") { setErrorMessege("ค่า ควันดำ ต้องมากกว่า 0 และไม่เกิน100"); }
        else { setErrorMessege("บันทึกไม่สำเร็จ"); }
    }

    const ValidateWheelCenter = (value: number) => {
        value > 0 ? setErrorWheelCenter(true) : setErrorWheelCenter(false);
    }

    const ValidateSoundLevel = (value: number) => {
        value > 0 ? setErrorSoundLevel(true) : setErrorSoundLevel(false);
    }

    const ValidateBlackSmoke = (value: number) => {
        value > 0 && value <= 100 ? setErrorBlackSmoke(true) : setErrorBlackSmoke(false);
    }

    return (
        <Page theme={pageTheme.service}>
            <Header
                title={`${profile.givenName}`}
            //subtitle="Some quick intro and links."
            >
                <div className={classes.margin}>
                    <Button
                        onClick={() => {
                            history.pushState("", "", "./carinspectionsearch");
                            window.location.reload(false);
                        }}
                        variant="contained"
                        color="primary"
                        style={{ width: 300, backgroundColor: "#34eb77", color: "#000000", fontSize: 18 }}
                        startIcon={<SearchIcon />}
                    >
                        <b>ค้นหาผลการตรวจสภาพรถ</b>
                    </Button>
                </div>
            </Header>
            <Content>
                <ContentHeader title="เพิ่มข้อมูลการตรวจสภาพรถ">
                    {status ? (
                        <div>
                            {alerttype != "" ? (
                                <Alert severity={alerttype} onClose={() => { setStatus(false) }}>
                                    {errormessege}
                                </Alert>
                            ) : null}
                        </div>
                    ) : null}
                </ContentHeader>

                <FormControl
                    className={classes.media}
                >
                    <CardMedia
                        component="img"
                        src={Image1Base64Function}
                    ></CardMedia>
                    <CardMedia
                        component="img"
                        src={Image2Base64Function}
                        style={{ marginTop: 30 }}
                    ></CardMedia>
                </FormControl>
                <div className={classes.root}>
                    <form noValidate autoComplete="off">

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
                                    value={users.filter((filter: EntUser) => filter.id == userid).map((item: EntUser) => `${item.name} (${item.email})`)}
                                    style={{ width: 400 }}
                                />
                            </FormControl>
                        </div>

                        <FormControl
                            className={classes.margin}
                            variant="outlined"
                        >
                            <InputLabel id="ambulance-label">รถพยาบาล</InputLabel>
                            <Select
                                labelId="ambulance-label"
                                id="ambulance"
                                value={ambulanceid}
                                onChange={AmbulancehandleChange}
                                style={{ width: 400 }}
                            >
                                {ambulances.filter((filter: any) => filter.edges?.Hasstatus?.resultName == "ส่งตรวจสภาพรถ").map((item: EntAmbulance) => (
                                    <MenuItem value={item.id}>{item.carregistration}</MenuItem>
                                ))}
                            </Select>
                        </FormControl>

                        <div>
                            <FormControl
                                className={classes.margin}
                                variant="outlined"
                            >
                                <TextField
                                    id="wheelcenter"
                                    label="ศูนย์ล้อ (เมตร)"
                                    type="number"
                                    size="medium"
                                    value={wheelcenterdata}
                                    helperText={errorWheelCenter ? "" : "ค่า ศูนย์ล้อ ต้องมากกว่า 0"}
                                    error={errorWheelCenter ? false : true}
                                    onChange={WheelCenterhandleChange}
                                    style={{ width: 100 }}
                                />
                            </FormControl>

                            <FormControl
                                className={classes.margin}
                                variant="outlined"
                            >
                                <TextField
                                    id="soundlevel"
                                    label="ระดับเสียง (เดซิเบล)"
                                    type="number"
                                    size="medium"
                                    value={soundleveldata}
                                    helperText={errorSoundLevel ? "" : "ค่า ระดับเสียง ต้องมากกว่า 0"}
                                    error={errorSoundLevel ? false : true}
                                    onChange={SoundLevelhandleChange}
                                    style={{ width: 100 }}
                                />
                            </FormControl>

                            <FormControl
                                className={classes.margin}
                                variant="outlined"
                            >
                                <TextField
                                    id="blacksmoke"
                                    label="ควันดำ (เปอร์เซ็นต์)"
                                    type="number"
                                    size="medium"
                                    value={blacksmokedata}
                                    helperText={errorBlackSmoke ? "" : "ค่า ควันดำ ต้องมากกว่า 0 และไม่เกิน100"}
                                    error={errorBlackSmoke ? false : true}
                                    onChange={BlackSmokehandleChange}
                                    style={{ width: 100 }}
                                />
                            </FormControl>
                        </div>

                        <div>
                            <FormControl
                                className={classes.margin}
                                variant="outlined"
                            >
                                <InputLabel id="result-label">ผลการตรวจสภาพ</InputLabel>
                                <Select
                                    labelId="result-label"
                                    id="result"
                                    value={inspectionresultid}
                                    onChange={InspectionResulthandleChange}
                                    style={{ width: 400 }}
                                >
                                    {inspectionresults.filter((filter: any) => filter.edges?.jobposition?.positionName == "เจ้าหน้าที่ตรวจสภาพรถ").map((item: EntInspectionResult) => (
                                        <MenuItem value={item.id}>{item.resultName}</MenuItem>
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
                                    id="note"
                                    label="หมายเหตุ"
                                    type="string"
                                    size="medium"
                                    value={note}
                                    onChange={NotehandleChange}
                                    style={{ width: 400 }}
                                />
                            </FormControl>
                        </div>

                        <div>
                            <FormControl
                                className={classes.margin}
                                variant="outlined"
                            >
                                <TextField
                                    id="datetime"
                                    label="เดือน/วัน/ปี"
                                    type="datetime-local"
                                    value={datetime}
                                    onChange={DateTimehandleChange}
                                    className={classes.textField}
                                    InputLabelProps={{
                                        shrink: true,
                                    }}
                                    style={{ width: 400 }}
                                />
                            </FormControl>
                        </div>

                        <div className={classes.margin}>
                            <Button
                                onClick={() => {
                                    CreateCarInspection();
                                }}
                                variant="contained"
                                color="primary"
                                style={{ width: 180, backgroundColor: "#365391" }}
                                startIcon={<SaveIcon />}
                            >
                                ยืนยัน
                            </Button>
                            <Button
                                onClick={() => {
                                    window.location.reload(false);
                                }}
                                variant="contained"
                                color="primary"
                                style={{ marginLeft: 40, width: 180, backgroundColor: "#1c0996" }}
                                startIcon={<CachedIcon />}
                            >
                                รีเฟรช
                            </Button>
                        </div>
                    </form>
                </div>
                <ComponanceTable></ComponanceTable>
            </Content>
        </Page>
    );
}
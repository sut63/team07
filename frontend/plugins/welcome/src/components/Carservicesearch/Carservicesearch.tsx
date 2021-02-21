import React, { useState, useEffect } from 'react';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
} from '@backstage/core';
import TextField from '@material-ui/core/TextField';
import { Link, Link as RouterLink } from 'react-router-dom';
import { Alert } from '@material-ui/lab';
import moment from 'moment';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import { EntCarservice } from '../../api/models/EntCarservice';
import { EntUser } from '../../api/models/EntUser';

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    root: {
      display: 'flex',
      flexWrap: 'wrap',
      justifyContent: 'center',
    },
    margin: {
      margin: theme.spacing(1),
    },

    withoutLabel: {
      marginTop: theme.spacing(1),
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
      width: 100,
    },
  }),
);

export default function Searchtable() {
  const classes = useStyles();
  const api = new DefaultApi();
  const [carservices, setCarservices] = useState([]);
  const [loading, setLoading] = useState(true);
  const profile = { givenName: 'ระบบบันทึกการใช้รถพยาบาล' };
  const [status, setStatus] = useState(false);
  const [alerttype, setAlertType] = useState(String);
  const [errormessege, setErrorMessege] = useState(String);
  const [users, setUsers] = useState<EntUser[]>([]);
  const [userid, setUser] = useState(Number);
  const [carservicesearch, setCarserviceSearch] = useState(String);

  useEffect(() => {

    const checkJobPosition = async () => {
      const jobdata = JSON.parse(String(localStorage.getItem("jobpositiondata")));
      setLoading(false);
      if (jobdata != "เจ้าหน้าที่โอเปอร์เรเตอร์") {
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



  const SearchCarservice = async () => {
    const apiUrl = `http://localhost:8080/api/v1/searchcarservices?customer=${carservicesearch}`;
    const requestOptions = {
      method: 'GET',
    };
    fetch(apiUrl, requestOptions)
      .then(response => response.json())
      .then(data => {
        console.log(data.data)
        setErrorMessege("ไม่พบข้อมูล");
        setAlertType("error");
        setCarservices([]);
        if (data.data != null) {
          if (data.data.length >= 1) {
            setErrorMessege("พบข้อมูล");
            setAlertType("success");
            console.log(data.data)
            setCarservices(data.data);
          }
        }

        setStatus(true);
      });
  }

  const handleSearchChange = (event: any) => {
    setCarserviceSearch(event.target.value as string);
  };

  return (
    <Page theme={pageTheme.library}>
      <Header
        title={`ท่านกำลังใช้งาน ${profile.givenName || ':)'}`}
        subtitle="สวัสดีครับท่านสมาชิกชมรมคนชอบรถพยาบาล"
      ><Link component={RouterLink} to="/Carservicemain">
          <Button variant="contained" color="primary" >
            ย้อนกลับ
     </Button>
        </Link>
      </Header>
      <Content>
        <ContentHeader title="ค้นหาบันทึกการใช้รถพยาบาล">
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

        <div className={classes.paper}><strong>ค้นหาชื่อผู้ใช้บริการ (เว้นว่าง เพื่อแสดงข้อมูลทั้งหมด)</strong></div>
        <TextField className={classes.textField}
          style={{ width: 400, marginLeft: 20, marginRight: -10 }}
          id="customer"
          label="กรอกชื่อผู้ใช้บริการ"
          variant="standard"
          color="secondary"
          type="string"
          size="medium"
          value={carservicesearch}
          onChange={handleSearchChange}
        />
        <div className={classes.margin}>
          <Button
            onClick={() => {
              SearchCarservice();
            }}
            variant="contained"
            color="primary"
          >
            ค้นหา
          </Button>
        </div>
        <TableContainer component={Paper}>
          <Table aria-label="simple table">
            <TableHead>
              <TableRow>
                <TableCell align="center">เลขที่</TableCell>
                <TableCell align="center">ชื่อผู้ใช้บริการ</TableCell>
                <TableCell align="center">อายุผู้ใช้บริการ</TableCell>
                <TableCell align="center">ที่อยู่</TableCell>
                <TableCell align="center">การใช้บริการ</TableCell>
                <TableCell align="center">ความเร่งด่วน</TableCell>
                <TableCell align="center">ระยะทาง</TableCell>
                <TableCell align="center">ผู้บันทึก</TableCell>
                <TableCell align="center">วัน-เวลา</TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {carservices.map((item : any) => (
                <TableRow key={item.id}>
                  <TableCell align="center">{item.id}</TableCell>
                  <TableCell align="center">{item.customer}</TableCell>
                  <TableCell align="center">{item.age}</TableCell>
                  <TableCell align="center">{item.location}</TableCell>
                  <TableCell align="center">{item.serviceinfo}</TableCell>
                  <TableCell align="center">{item.edges?.Urgentid?.urgent}</TableCell>
                  <TableCell align="center">{item.edges?.Disid?.Distance}</TableCell>
                  <TableCell align="center">{item.edges?.Userid?.name}</TableCell>
                  <TableCell align="center">{moment(item.Datetime).format('DD/MM/YYYY HH.mm น.')}</TableCell>
                  <TableCell align="center">
                  </TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
        <br></br>
      </Content>
    </Page>
  );
};

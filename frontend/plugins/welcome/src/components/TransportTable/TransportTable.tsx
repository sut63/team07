import React, { useState, useEffect,FC } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import { EntTransport } from '../../api/models/EntTransport';
import { Link as RouterLink } from 'react-router-dom';
import {
  Content,
  Header,
  Page,
  pageTheme,
  ContentHeader,
  Link,
 } from '@backstage/core';
const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});
 
export default function ComponentsTable() {
 const classes = useStyles();
 const api = new DefaultApi();
 const [transports, setTransports] = useState<EntTransport[]>([]);
 const [loading, setLoading] = useState(true);
 
 useEffect(() => {
   const getTransports = async () => {
     const res = await api.listTransport();
     setLoading(false);
     setTransports(res);
   };
   getTransports();

 }, [loading]);
 
 const deleteTransports = async (id: number) => {
   const res = await api.deleteTransport({ id: id });
   setLoading(true);
 };
 
  const profile = { givenName: 'ข้อมูลการลงทะเบียนรถ' };
  const HeaderCustom = {
      minHeight: '50px',
    };
 return (

  <Page theme={pageTheme.service} >
  &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
<Header
 style={HeaderCustom}
 title={`${profile.givenName}`}
 subtitle=""
>
<Link component={RouterLink} to="/CreateTransport">
            <Button variant="contained" color="primary" style={{backgroundColor: "#9834AE"}}>
              กลับสู่หน้าส่งตัวผู้ป่วย
            </Button>
          </Link>
</Header>
   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
       
         <TableRow>
           <TableCell align="center">เลขที่</TableCell>
           <TableCell align="center">เจ้าหน้าที่</TableCell>
           <TableCell align="center">รถพยาบาล</TableCell>
           <TableCell align="center">โรงพยาบาลต้นทาง</TableCell>
           <TableCell align="center">โรงพยาบาลปลายทาง</TableCell>
           <TableCell align="center">อาการ</TableCell>
           <TableCell align="center">การแพ้ยา</TableCell>
           <TableCell align="center">หมายเหตุเพิ่มเติม</TableCell>
         </TableRow>
       </TableHead>
       <TableBody>
         {transports.map((item:any) => (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.edges.user.name}</TableCell>
             <TableCell align="center">{item.edges.ambulance.carregistration}</TableCell>
             <TableCell align="center">{item.edges?.send?.hospital}</TableCell>
             <TableCell align="center">{item.edges?.receive?.hospital}</TableCell>
             <TableCell align="center">{item.symptom}</TableCell>
             <TableCell align="center">{item.drugallergy}</TableCell>
             <TableCell align="center">{item.note}</TableCell>
             <TableCell align="center">
               <Button
                 onClick={() => {
                  deleteTransports(item.id);
                 }}
                 style={{ marginLeft: 10 , backgroundColor: "#140087" }}
                 variant="contained"
                 color="secondary"
               >
                 ลบข้อมูล
               </Button>
             </TableCell>
           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>
   </Page>
 );
}

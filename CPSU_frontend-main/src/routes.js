// AppRoutes.js
import { Routes, Route, Navigate } from "react-router-dom";
import News from "./page/news";
import NewsDetail from "./page/news_detail";
import Admission from "./page/admission";
import Personnel from "./page/personnel";
import Course from "./page/course";
import Subject from "./page/subject";

import Add_News from "./page/admin/news_add";
import Detail_News from "./page/admin/news_detail";
import Edit_News from "./page/admin/news_edit";

import Course_admin from "./page/admin/course";
import Add_course from "./page/admin/course_add";
import Edit_course from "./page/admin/course_edit";
import Detail_course from "./page/admin/course_detail";

import Subject_admin from "./page/admin/subject";
import Add_Subject from "./page/admin/subject_add";
import Edit_Subject from "./page/admin/subject_edit";
import Detail_Subject from "./page/admin/subject_detail";

import Personnel_admin from "./page/admin/personnel";
import Add_Personnel from "./page/admin/personnel_add";
import Edit_Personnel from "./page/admin/personnel_edit";
import Detail_Personnel from "./page/admin/personnel_detail";

import TeacherDetail from "./page/personnel_teacher_detail";
import StaffDetail from "./page/personnel_staff_detail";
import Homepage from "./page/homepage";
import News_admin from "./page/admin/news";
import About from "./page/about";
import Login from "./component/login";
import Course_detail from "./page/course_detail";
import Document from "./page/document";
import EventCalendarAdmin from "./page/admin/EventCalendarAdmin";
import UserPermissions from "./page/admin/UserPermissions";
import Subject_detail from "./page/subject_detail";

import EditTeacherInfomation from "./page/teacher/editinfomation";
import TeacherInfomation from "./page/teacher/infomation";

import Dashboard from "./page/admin/dashboard";
import EventFormAdd from "./page/admin/EventFormAdd";
import EventCalendarPublic from "./page/EventCalendarPublic";

const AppRoutes = () => {
  return (
    <Routes>
      {/* Default Redirect */}
      <Route path="/" element={<Navigate to="/home" />} />
      {/* Admin Group */}

      <Route path="/admin" >
        <Route path="addnews" element={<Add_News />} />
        <Route path="news" element={<News_admin />} />
        <Route path="news/:id" element={<Detail_News />} />
        <Route path="editnews/:id" element={<Edit_News />} />

        <Route path="addcourse" element={<Add_course />} />
        <Route path="course" element={<Course_admin />} />
        <Route path="course/:id" element={<Detail_course />} />
        <Route path="editcourse/:id" element={<Edit_course />} />

        <Route path="addsubject" element={<Add_Subject />} />
        <Route path="subject" element={<Subject_admin />} />
        <Route path="subject/:id" element={<Detail_Subject />} />
        <Route path="editsubject/:id" element={<Edit_Subject />} />

        <Route path="addpersonnel" element={<Add_Personnel />} />
        <Route path="personnel" element={<Personnel_admin />} />
        <Route path="personnel/:id" element={<Detail_Personnel />} />
        <Route path="editpersonnel/:id" element={<Edit_Personnel />} />

        <Route path="calendar" element={<EventCalendarAdmin />} />
        <Route path="calendar/add" element={<EventFormAdd />} />

        <Route path="userpermissions" element={<UserPermissions />} />

        <Route path="dashboard" element={<Dashboard />} />
      </Route>

      <Route path="/teacher">
        <Route path="information" element={<TeacherInfomation />} />
        <Route path="informationedit" element={<EditTeacherInfomation />} />
      </Route>
      {/* User Group */}
      {/* <Route path="/user" element={<UserLayout />}>
        <Route path="home" element={<UserHome />} />
        <Route path="news" element={<UserNews />} />
        <Route path="profile" element={<UserProfile />} />
      </Route> */}
      <Route path="/login" element={<Login />}></Route>
      <Route path="/home" element={<Homepage />}></Route>
      <Route path="/admission" element={<Admission />} />
      <Route path="/personnel" element={<Personnel />} />
      <Route path="/course" element={<Course />} />
      <Route path="/course/:course_id" element={<Course_detail />} />
      <Route path="/subject" element={<Subject />} />
      <Route path="/news" element={<News />} />
      <Route path="/news/:news_id" element={<NewsDetail />} />
      <Route path="/about" element={<About />} />
      <Route path="/document" element={<Document />} />
      <Route path="/staff/:id" element={<StaffDetail />} />
      <Route path="/teacher/:id" element={<TeacherDetail />} />
      <Route path="/subject/:subject_id" element={<Subject_detail />} />
      <Route path="/calendar" element={<EventCalendarPublic />} />

    </Routes>
  );
};

export default AppRoutes;

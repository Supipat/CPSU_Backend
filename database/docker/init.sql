-- create news

CREATE TABLE IF NOT EXISTS news_types (
    type_id SERIAL PRIMARY KEY,
    type_name VARCHAR(50) NOT NULL
);

CREATE TABLE IF NOT EXISTS news (
    news_id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    type_id INT NOT NULL,
    detail_url TEXT NOT NULL,
    cover_image TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (type_id) REFERENCES news_types(type_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS news_images (
    image_id SERIAL PRIMARY KEY,
    news_id INT NOT NULL,
    file_image TEXT NOT NULL,
    FOREIGN KEY (news_id) REFERENCES news(news_id) ON DELETE CASCADE
);

-- create courses

CREATE TABLE IF NOT EXISTS degree (
    degree_id SERIAL PRIMARY KEY,
    degree_name VARCHAR(20) NOT NULL
);

CREATE TABLE IF NOT EXISTS majors (
    major_id SERIAL PRIMARY KEY,
    major_name VARCHAR(20) NOT NULL
);

CREATE TABLE IF NOT EXISTS courses (
    course_id SERIAL PRIMARY KEY,
    degree_id INT NOT NULL,
    major_id INT NOT NULL,
    year INT NOT NULL,
    thai_course VARCHAR(255) NOT NULL,
    eng_course VARCHAR(255) NOT NULL,
    thai_degree VARCHAR(255) NOT NULL,
    eng_degree VARCHAR(255) NOT NULL,
    admission_req TEXT NOT NULL,
    graduation_req TEXT NOT NULL,
    philosophy TEXT NOT NULL,
    objective TEXT NOT NULL,
    tuition TEXT NOT NULL,
    credits VARCHAR(255) NOT NULL,
    career_paths TEXT NOT NULL,
    plo TEXT NOT NULL,
    detail_url TEXT NOT NULL,
    FOREIGN KEY (degree_id) REFERENCES degree(degree_id) ON DELETE CASCADE,
    FOREIGN KEY (major_id) REFERENCES majors(major_id) ON DELETE CASCADE
);

-- create roadmap

CREATE TABLE IF NOT EXISTS roadmap (
    roadmap_id SERIAL PRIMARY KEY,
    course_id INT NOT NULL,
    roadmap_url TEXT NOT NULL,
    FOREIGN KEY (course_id) REFERENCES courses(course_id) ON DELETE CASCADE
);

-- TRIGGER

CREATE OR REPLACE FUNCTION update_modified_column()
RETURNS TRIGGER AS $$
BEGIN
    NEW.updated_at = now(); 
    RETURN NEW;
END;
$$ LANGUAGE 'plpgsql';

CREATE TRIGGER update_news_modtime
BEFORE UPDATE ON news
FOR EACH ROW
EXECUTE PROCEDURE update_modified_column();

-- insert news

INSERT INTO news_types(type_name) VALUES
('ข่าวประชาสัมพันธ์'),
('ทุนการศึกษา'),
('รางวัลที่ได้รับ'),
('กิจกรรมของภาควิชา');

INSERT INTO news(title,content,type_id,detail_url,cover_image) VALUES
('คู่มือแนะนำนักศึกษาใหม่ ปีการศึกษา 2568',
    'คู่มือแนะนำ การลงทะเบียนรายวิชาเรียน การเพิ่ม - ถอน - การเปลี่ยนกลุ่มเรียน ในรายวิชาเดิม การถอนรายวิชาโดยติดสัญลักษณ์ W ช่องทางการชำระค่าธรรมเนียมการศึกษา เอกสารแนบเบิกค่าเล่าเรียน การลาพักการศึกษา นักศึกษาที่คาดว่าจะสำเร็จการศึกษา การขอหนังสือสำคัญทางการศึกษา การติดต่อเจ้าหน้าที่งานทะเบียนและประมวลผล ช่องทางติดต่อคณะวิชา กองกิจการนักศึกษา การให้บริการสำหรับนักศึกษา วิทยาเขตสารสนเทศเพชรบุรี สำนักดิจิทัลเทคโนโลยี หอสมุด มหาวิทยาลัยศิลปากร ประกันอุบัติเหตุส่วนบุคคล ศูนย์บริการสุขภาพ คณะเภสัชศาสตร์ บริการด้านสุขภาพร่างกาย เครื่องแบบนักศึกษา / บัตรนักศึกษาอิเล็กทรอนิกส์',
    (SELECT type_id FROM news_types WHERE type_name = 'ข่าวประชาสัมพันธ์' LIMIT 1),
    'https://drive.google.com/file/d/1FnJRketlluku27HQqs-mVPnxMJ5wFxqZ/view?usp=sharing&fbclid=IwZXh0bgNhZW0CMTAAAR3M6TbS4tN1DUSa-z2NaM1ekAOELZp7TnsIhJWC6g_dvfz-sD_b0La0S7U_aem_K5_Zhi2eLKhIpJxaeszOlQ',
    'https://cpsu-website.s3.ap-southeast-2.amazonaws.com/images/news/manual.jpg'),
('เปิดรับสมัครปริญญาโทและปริญญาเอกสาขาเทคโนโลยีสารสนเทศและนวัตกรรมดิจิทัล',
    'เปิดรับสมัครแล้ว! ภาควิชาคอมพิวเตอร์ คณะวิทยาศาสตร์ มหาวิทยาลัยศิลปากร เปิดรับสมัครนักศึกษาระดับบัณฑิตศึกษา ปริญญาโท - ปริญญาเอก หลักสูตร IT สำหรับคนทำงาน สู่ผู้เชี่ยวชาญด้านวิจัยและนวัตกรรม เน้นเรียนรู้กระบวนทำวิจัยเพื่อแก้ปัญหาจริง มีทุนผู้ช่วยวิจัยและทุนนำเสนอผลงานในงานประชุมวิชาการ มีทั้งภาคปกติ (เรียนวันธรรมดา) และโครงการพิเศษ (เรียน ส อา) มี 3 แผนการเรียนให้เลือก Data Science, Project Management, DevOps เปิดรับสมัครรอบ 1 วันที่ 16 ธ.ค. 67 - 21 ก.พ. 68 เปิดรับสมัครรอบ 2 วันี่ต่ 3 มี.ค. 68 - 7 พ.ค. 68 สมัครง่าย ๆ ผ่านช่องทางออนไลน์ได้ที่ https://graduate.su.ac.th มีข้อสงสัยสอบถามเพิ่มเติมได้ที่ เพจ Facebook : https://www.facebook.com/computingsu/ หรือ โทร 034-272-923',
    (SELECT type_id FROM news_types WHERE type_name = 'ข่าวประชาสัมพันธ์' LIMIT 1),
    'https://graduate.su.ac.th',
    'https://cpsu-website.s3.ap-southeast-2.amazonaws.com/images/news/degree1.jpg'),
('กำหนดการรับสมัครและสัมภาษณ์ทุนการศึกษา ประจำปีการศึกษา 2568 (ครั้งที่ 1)',
    'กำหนดการรับสมัครและสัมภาษณ์ทุนการศึกษา ของคณะวิทยาศาสตร์ เพื่อขอทุนการศึกษาสำหรับนักศึกษาปริญญาตรี ภาคการศึกษาต้น ปีการศึกษา 2568 สังกัดคณะวิทย์ ขาดแคลนทุน เกรดไม่ต่ำกว่า 2.00 ขอและยื่นใบสัมครได้ที่ " งานบริการการศึกษา ชั้น 1 อาคารวิทย์ 1 "',
    (SELECT type_id FROM news_types WHERE type_name = 'ทุนการศึกษา' LIMIT 1),
    'https://dsa.su.ac.th/ksu/wp-content/uploads/2025/06/%E0%B8%9B%E0%B8%A3%E0%B8%B0%E0%B8%81%E0%B8%B2%E0%B8%A8%E0%B8%81%E0%B8%AD%E0%B8%87%E0%B8%81%E0%B8%B4%E0%B8%88%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B8%99%E0%B8%B1%E0%B8%81%E0%B8%A8%E0%B8%B6%E0%B8%81%E0%B8%A9%E0%B8%B2-%E0%B9%80%E0%B8%A3%E0%B8%B7%E0%B9%88%E0%B8%AD%E0%B8%87-%E0%B8%81%E0%B8%B3%E0%B8%AB%E0%B8%99%E0%B8%94%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B8%A3.pdf',
    'https://cpsu-website.s3.ap-southeast-2.amazonaws.com/images/news/scholarship.jpg'),
('ขอแสดงความยินดีกับ ผู้ช่วยศาสตราจารย์ ดร.อรวรรณ เชาวลิต อาจารย์ประจำภาควิชาคอมพิวเตอร์',
    'คณะวิทยาศาสตร์ มหาวิทยาลัยศิลปากร ขอแสดงความยินดีกับ ผู้ช่วยศาสตราจารย์ ดร.อรวรรณ เชาวลิต ภาควิชาคอมพิวเตอร์ ในโอกาสที่ตีพิมพ์ผลงานวิจัยในวารสาร ICIC Express Letters, Part B: Applications ในฐานข้อมูล Scopus (ScimagoJR, Quartile 4) เรื่อง Multivariate time series forecasting Thailand soybean meal price with deep learning models ICIC Express Letters, Part B: Applications 2025, 16(3), 317-323. http://www.icicelb.org/ellb/contents/2025/3/elb-16-03-10.pdf SDG 8 DECENT WORK AND ECONOMIC GROWTH',
    (SELECT type_id FROM news_types WHERE type_name = 'รางวัลที่ได้รับ' LIMIT 1),
    'http://www.icicelb.org/ellb/contents/2025/3/elb-16-03-10.pdf',
    'https://cpsu-website.s3.ap-southeast-2.amazonaws.com/images/news/reward-orawan.jpg'),
('ขอแสดงความยินดีกับผู้ช่วยศาสตราจารย์ ดร.ปัญญนัท อ้นพงษ์ เนื่องในโอกาสได้รับการแต่งตั้งให้ดำรงตำแหน่งทางวิชาการ',
    'ภาควิชาคอมพิวเตอร์ คณะวิทยาศาสตร์ มหาวิทยาลัยศิลปากร ขอแสดงความยินดีกับผู้ช่วยศาสตราจารย์ ดร.ปัญญนัท อ้นพงษ์ เนื่องในโอกาสได้รับการแต่งตั้งให้ดำรงตำแหน่งทางวิชาการ ตำแหน่ง “ผู้ช่วยศาสตราจารย์” ในสาขาวิชาเทคโนโลยีสารสนเทศ (1806)',
    (SELECT type_id FROM news_types WHERE type_name = 'รางวัลที่ได้รับ' LIMIT 1),
    'https://www.facebook.com/ScienceSilpakornUniversity/posts/คณะวิทยาศาสตร์-มหาวิทยาลัยศิลปากร-ขอแสดงความยินดีกับผู้ช่วยศาสตราจารย์-ดรปัญญนัท/982826210555894/',
    'https://cpsu-website.s3.ap-southeast-2.amazonaws.com/images/news/reward-panyanut.jpg'),
('ขอแสดงความยินดีกับ นายภากร กัทชลี นักศึกษา ปริญญาเอก หลักสูตรเทคโนโลยีสารสนเทศและนวัตกรรมดิจิทัล',
    'ภาควิชาคอมพิวเตอร์ ขอแสดงความยินดีกับ นายภากร กัทชลี นักศึกษา ปริญญาเอก หลักสูตรเทคโนโลยีสารสนเทศและนวัตกรรมดิจิทัล ที่ได้รับรางวัล Best Presentation ในงาน The 6th Asia Joint Conference on Computing (AJCC 2025) ณ เมือง Osaka ประเทศญี่ปุ่นจาก Paper หัวข้อ "Intention Classification of Chinese Topics on Thai Facebook Pages Using Transformer Models with Emotional Features" อาจารย์ที่ปรึกษา อ.ดร.สัจจาภรณ์ ไวจรรยา และ ผศ.ดร.ณัฐโชติ พรหมฤทธิ์',
    (SELECT type_id FROM news_types WHERE type_name = 'รางวัลที่ได้รับ' LIMIT 1),
    'https://www.facebook.com/computingsu/posts/-ภาควิชาคอมพิวเตอร์-ขอแสดงความยินดีกับ-นายภากร-กัทชลี-นักศึกษา-ปริญญาเอก-หลักสูต/122207080106177331/',
    'https://cpsu-website.s3.ap-southeast-2.amazonaws.com/images/news/reward-phakon1.jpg'),
('ขอแสดงความยินดีกับ  นายณัฐพิสิษ กังวานธรรมกุล,นายนัทธพงศ์ เป็กทองและนาย พิรพัฒน์ ยิ่งแก้ว นักศึกษาสาขาวิชาเทคโนโลยีสารสนเทศ กับผลงาน "การพัฒนาระบบอํานวยการสอบ คณะวิทยาศาสตร์ มหาวิทยาลัยศิลปากร',
    'ภาควิชาคอมพิวเตอร์ ขอแสดงความยินดีกับนักศึกษาและอาจารย์ที่ปรึกษา ที่เข้าร่วมนำเสนอผลงานและได้รับรางวัลในงาน การประชุมวิชาการระดับปริญญาตรีด้านคอมพิวเตอร์ภูมิภาคเอเชีย ครั้งที่ 13 (The 13th Asia Undergraduate Conference on Computing: AUCC 2025) ในระหว่างวันที่ 19 - 21 กุมภาพันธ์ พ.ศ. 2568 ณ มหาวิทยาลัยราชภัฏอุตรดิตถ์ จังหวัดอุตรดิตถ์ รางวัลรองชนะเลิศอันดับ 1 Track นวัตกรรม กับผลงาน "การพัฒนาระบบอํานวยการสอบ คณะวิทยาศาสตร์ มหาวิทยาลัยศิลปากร (Development of examination Management system for the Faculty of Science)" ซึ่งผลงานนี้ได้มีการนำไปใช้จริงทดแทนระบบเดิม ในการจัดสอบของคณะฯ เมื่อช่วงสอบกลางภาคการศึกษาที่ผ่านมา นักศึกษาที่เข้านำเสนอ - นาย ณัฐพิสิษ กังวานธรรมกุล - นาย นัทธพงศ์ เป็กทอง - นาย พิรพัฒน์ ยิ่งแก้ว นักศึกษาสาขาวิชาเทคโนโลยีสารสนเทศ อาจารย์ที่ปรึกษา ผศ.ดร.ณัฐโชติ พรหมฤทธิ์ และ อ.ดร.สัจจาภรณ์ ไวจรรยา',
    (SELECT type_id FROM news_types WHERE type_name = 'รางวัลที่ได้รับ' LIMIT 1),
    'https://www.facebook.com/computingsu/posts/-ภาควิชาคอมพิวเตอร์-ขอแสดงความยินดีกับนักศึกษาและอาจารย์ที่ปรึกษา-ที่เข้าร่วมนำเ/122194936622177331/',
    'https://cpsu-website.s3.ap-southeast-2.amazonaws.com/images/news/reward2nd1.jpg'),
('งานสานสัมพันธ์ภาคคอมพิวเตอร์',
    'เตรียมตัวให้พร้อม! งานสานสัมพันธ์ภาคคอมพิวเตอร์กำลังจะเริ่มขึ้น! เจอกันวันที่ 28 กุมภาพันธ์ 2568 ณ ลานจอดรถตึก 4 เวลา 17:30 - 21:30 น. ธีม: ย้อนยุค ไม่จำกัดช่วงเวลา! จะเป็นยุคหิน มนุษย์ถ้ำ อารยธรรมโบราณ หรือยุคใดก็ได้ จัดเต็มมาให้สุด! มีรางวัลแต่งกายยอดเยี่ยม 3 รางวัล!ใครแต่งตัวเข้าธีมและโดดเด่นที่สุด มีสิทธิ์คว้ารางวัลไปเลย! สนุกไปกับการแสดงสุดพิเศษ, เกมสุดมันส์, ลุ้นรางวัล และดนตรีสดปิดท้าย! ใครลงชื่อไว้แล้ว เจอกันแน่นอน! เตรียมชุดให้พร้อม แล้วมาสนุกไปด้วยกัน!',
    (SELECT type_id FROM news_types WHERE type_name = 'กิจกรรมของภาควิชา' LIMIT 1),
    '',
    'https://cpsu-website.s3.ap-southeast-2.amazonaws.com/images/news/cpsu_event.png');

INSERT INTO news_images(news_id,file_image) VALUES
((SELECT news_id FROM news WHERE title = 'คู่มือแนะนำนักศึกษาใหม่ ปีการศึกษา 2568' LIMIT 1),
    'https://cpsu-website.s3.ap-southeast-2.amazonaws.com/images/news/manual.jpg'),
((SELECT news_id FROM news WHERE title = 'เปิดรับสมัครปริญญาโทและปริญญาเอกสาขาเทคโนโลยีสารสนเทศและนวัตกรรมดิจิทัล' LIMIT 1),
    'https://cpsu-website.s3.ap-southeast-2.amazonaws.com/images/news/degree1.jpg'),
((SELECT news_id FROM news WHERE title = 'เปิดรับสมัครปริญญาโทและปริญญาเอกสาขาเทคโนโลยีสารสนเทศและนวัตกรรมดิจิทัล' LIMIT 1),
    'https://cpsu-website.s3.ap-southeast-2.amazonaws.com/images/news/degree2.jpg'),
((SELECT news_id FROM news WHERE title = 'เปิดรับสมัครปริญญาโทและปริญญาเอกสาขาเทคโนโลยีสารสนเทศและนวัตกรรมดิจิทัล' LIMIT 1),
    'https://cpsu-website.s3.ap-southeast-2.amazonaws.com/images/news/degree3.jpg'),
((SELECT news_id FROM news WHERE title = 'เปิดรับสมัครปริญญาโทและปริญญาเอกสาขาเทคโนโลยีสารสนเทศและนวัตกรรมดิจิทัล' LIMIT 1),
    'https://cpsu-website.s3.ap-southeast-2.amazonaws.com/images/news/degree4.jpg'),
((SELECT news_id FROM news WHERE title = 'กำหนดการรับสมัครและสัมภาษณ์ทุนการศึกษา ประจำปีการศึกษา 2568 (ครั้งที่ 1)' LIMIT 1),
    'https://cpsu-website.s3.ap-southeast-2.amazonaws.com/images/news/scholarship.jpg'),
((SELECT news_id FROM news WHERE title = 'ขอแสดงความยินดีกับ ผู้ช่วยศาสตราจารย์ ดร.อรวรรณ เชาวลิต อาจารย์ประจำภาควิชาคอมพิวเตอร์' LIMIT 1),
    'https://cpsu-website.s3.ap-southeast-2.amazonaws.com/images/news/reward-orawan.jpg'),
((SELECT news_id FROM news WHERE title = 'ขอแสดงความยินดีกับผู้ช่วยศาสตราจารย์ ดร.ปัญญนัท อ้นพงษ์ เนื่องในโอกาสได้รับการแต่งตั้งให้ดำรงตำแหน่งทางวิชาการ' LIMIT 1),
    'https://cpsu-website.s3.ap-southeast-2.amazonaws.com/images/news/reward-panyanut.jpg'),
((SELECT news_id FROM news WHERE title = 'ขอแสดงความยินดีกับ นายภากร กัทชลี นักศึกษา ปริญญาเอก หลักสูตรเทคโนโลยีสารสนเทศและนวัตกรรมดิจิทัล' LIMIT 1),
    'https://cpsu-website.s3.ap-southeast-2.amazonaws.com/images/news/reward-phakon1.jpg'),
((SELECT news_id FROM news WHERE title = 'ขอแสดงความยินดีกับ นายภากร กัทชลี นักศึกษา ปริญญาเอก หลักสูตรเทคโนโลยีสารสนเทศและนวัตกรรมดิจิทัล' LIMIT 1),
    'https://cpsu-website.s3.ap-southeast-2.amazonaws.com/images/news/reward-phakon2.jpg'),
((SELECT news_id FROM news WHERE title = 'ขอแสดงความยินดีกับ  นายณัฐพิสิษ กังวานธรรมกุล,นายนัทธพงศ์ เป็กทองและนาย พิรพัฒน์ ยิ่งแก้ว นักศึกษาสาขาวิชาเทคโนโลยีสารสนเทศ กับผลงาน "การพัฒนาระบบอํานวยการสอบ คณะวิทยาศาสตร์ มหาวิทยาลัยศิลปากร' LIMIT 1),
    'https://cpsu-website.s3.ap-southeast-2.amazonaws.com/images/news/reward2nd1.jpg'),
((SELECT news_id FROM news WHERE title = 'ขอแสดงความยินดีกับ  นายณัฐพิสิษ กังวานธรรมกุล,นายนัทธพงศ์ เป็กทองและนาย พิรพัฒน์ ยิ่งแก้ว นักศึกษาสาขาวิชาเทคโนโลยีสารสนเทศ กับผลงาน "การพัฒนาระบบอํานวยการสอบ คณะวิทยาศาสตร์ มหาวิทยาลัยศิลปากร' LIMIT 1),
    'https://cpsu-website.s3.ap-southeast-2.amazonaws.com/images/news/reward2nd2.jpg'),
((SELECT news_id FROM news WHERE title = 'ขอแสดงความยินดีกับ  นายณัฐพิสิษ กังวานธรรมกุล,นายนัทธพงศ์ เป็กทองและนาย พิรพัฒน์ ยิ่งแก้ว นักศึกษาสาขาวิชาเทคโนโลยีสารสนเทศ กับผลงาน "การพัฒนาระบบอํานวยการสอบ คณะวิทยาศาสตร์ มหาวิทยาลัยศิลปากร' LIMIT 1),
    'https://cpsu-website.s3.ap-southeast-2.amazonaws.com/images/news/reward2nd3.jpg'),
((SELECT news_id FROM news WHERE title = 'ขอแสดงความยินดีกับ  นายณัฐพิสิษ กังวานธรรมกุล,นายนัทธพงศ์ เป็กทองและนาย พิรพัฒน์ ยิ่งแก้ว นักศึกษาสาขาวิชาเทคโนโลยีสารสนเทศ กับผลงาน "การพัฒนาระบบอํานวยการสอบ คณะวิทยาศาสตร์ มหาวิทยาลัยศิลปากร' LIMIT 1),
    'https://cpsu-website.s3.ap-southeast-2.amazonaws.com/images/news/reward2nd4.jpg'),
((SELECT news_id FROM news WHERE title = 'งานสานสัมพันธ์ภาคคอมพิวเตอร์' LIMIT 1),
    'https://cpsu-website.s3.ap-southeast-2.amazonaws.com/images/news/cpsu_event.png');

-- insert courses

INSERT INTO degree(degree_name) VALUES
('ปริญญาตรี'),('ปริญญาโท'),('ปริญญาเอก');

INSERT INTO majors(major_name) VALUES
('วิทยาการคอมพิวเตอร์'),('เทคโนโลยีสารสนเทศ'),('วิทยาการข้อมูล');

INSERT INTO courses(degree_id,major_id,year,thai_course,eng_course,thai_degree,eng_degree,admission_req,graduation_req,philosophy,objective,tuition,credits,career_paths,plo,detail_url) VALUES
((SELECT degree_id FROM degree WHERE degree_name = 'ปริญญาตรี' LIMIT 1),
    (SELECT major_id FROM majors WHERE major_name = 'เทคโนโลยีสารสนเทศ' LIMIT 1),
    2565,'(วท.บ) หลักสูตรวิทยาศาสตรบัณฑิต สาขาวิชาเทคโนโลยีสารสนเทศ 2565','Bachelor of Science Program in Information Technology 2022',
    'วิทยาศาสตรบัณฑิต  (เทคโนโลยีสารสนเทศ)','Bachelor of Science (Information Technology)','สำเร็จการศึกษาระดับมัธยมศึกษาปีที่ 6 หรือเทียบเท่า','เกรดเฉลี่ยไม่ต่ำกว่า 2.00 เกรดเฉลี่ยวิชาเฉพาะไม่ต่ำกว่า 2.00',
    'สร้างบัณฑิตที่มีคุณภาพ คุณธรรม จริยธรรม และวินัย มีความรู้และทักษะทางด้านเทคโนโลยีสารสนเทศที่ทันสมัย มีความคิดสร้างสรรค์ สามารถบูรณาการความรู้ไปประยุกต์ใช้งานด้านต่างๆ สอดคล้องกับความต้องการทางด้านเทคโนโลยีสารสนเทศของทั้งภาครัฐและเอกชน',
    'หลักสูตรนี้มุ่งผลิตบัณฑิตที่มีคุณภาพมีความรู้ความเชี่ยวชาญในสาขาเทคโนโลยีสารสนเทศที่ทันสมัย โดยมุ่งเน้นให้บัณฑิตสามารถหาความรู้ด้าน เทคโนโลยีสารสนเทศไปประยุกต์ใช้ในการทำงาน และเป็นพื้นฐานในการศึกษาต่อในสาขาวิชาที่เกี่ยวข้องต่อไปในอนาคตนอกจากนี้ยังมุ่งเน้นให้ นักศึกษามองเห็นถึงความสัมพันธ์ของศาสตร์ต่างๆในสาขาวิชาเทคโนโลยีสารสนเทศนำความรู้เหล่านั้นมาสร้างงานประยุกต์เพื่อเตรียมพร้อม ในการทำงาน และการวิจัยในขั้นสูงต่อไป',
    'ประมาณ 20,000 บาทต่อเทอม','จำนวนไม่น้อยกว่า 133 หน่วยกิต',
    '1) นักเทคโนโลยีสารสนเทศ หรือนักเทคโนโลยีและสารสนเทศ 2) นักพัฒนาระบบ นักพัฒนาเว็บไซต์ 3) ผู้ดูแลระบบฐานข้อมูล 4) นักวิเคราะห์ และออกแบบระบบงานสารสนเทศ 5) ผู้ดูแลระบบเครือข่าย และเครื่องแม่ข่าย 6) ผู้จัดการโครงการสารสนเทศ 7) ผู้จัดการซอฟต์แวร์ หรือผู้จัดการเทคโนโลยีสารสนเทศ 8) นักทดสอบระบบในสถานประกอบการที่มีการใช้เทคโนโลยีสารสนเทศ 9) นักวิเคราะห์ข้อมูลทางธุรกิจ 10) นักวิทยาศาสตร์ข้อมูล',
    'PLO1	อธิบายความหมายและคุณค่าของศิลปะและการสร้างสรรค์ได้ PLO2	อภิปรายความหมายของความหลากหลายทางวัฒนธรรมได้ PLO3	ระบุความรู้เบื้องต้นเกี่ยวกับการประกอบธุรกิจและทักษะพื้นฐานที่จำเป็นต่อการเป็นผู้ประกอบการได้ PLO4	มีทักษะการใช้ภาษา และสื่อสารได้ตรงตามวัตถุประสงค์ในบริบทการสื่อสารที่หลากหลาย PLO5	เลือกใช้เทคโนโลยีสารสนเทศและการสื่อสารได้ตรงตามวัตถุประสงค์ ตลอดจนรู้เท่าทันสื่อและสารสนเทศ PLO6	แสวงหาความรู้ได้ด้วยตนเอง และนำความรู้ไปใช้ในการพัฒนาตนเองและการดำเนินชีวิต PLO7	แสดงออกซึ่งทักษะความสัมพันธ์ระหว่างบุคคล สามารถทำงานร่วมกับผู้อื่นได้ มีระเบียบวินัย ตรงต่อเวลา ซื่อสัตย์สุจริต มีความรับผิดชอบต่อตนเอง สังคม และสิ่งแวดล้อม PLO8	ใช้ความคิดสร้างสรรค์ในการสร้างผลงานหรือดำเนินโครงการได้ PLO9	คิดวิเคราะห์ วางแผน อย่างเป็นระบบ เพื่อแก้ไขปัญหาหรือเพื่อออกแบบนวัตกรรมได้ PLO10	อธิบายหลักการและองค์ประกอบของเทคโนโลยีสารสนเทศได้ PLO11	อธิบายสาระสำคัญของจริยธรรมและกฎหมายทางด้านเทคโนโลยีสารสนเทศ PLO12	ออกแบบ ติดตั้ง และจัดการระบบฐานข้อมูลได้ PLO13	ประยุกต์ใช้หลักการของเครือข่ายคอมพิวเตอร์และกลไกสำหรับรักษาความปลอดภัยของระบบสารสนเทศได้ PLO14	พัฒนาระบบเว็บแอปพลิเคชันให้เหมาะสมกับงานทางธุรกิจได้ PLO15	จัดเตรียมสภาวะแวดล้อมที่เหมาะสมต่อการพัฒนาระบบสารสนเทศได้ PLO16	ติดตั้ง ทดสอบ และบำรุงรักษา ระบบสารสนเทศที่พัฒนาขึ้นได้ PLO17	เก็บรวบรวมข้อมูล จัดการข้อมูล วิเคราะห์ข้อมูล และนำเสนอข้อมูลในรูปแบบที่หลากหลายได้ PLO18	พัฒนาโปรแกรมประยุกต์ได้ PLO19	รวบรวม สืบค้น ทดลองประยุกต์ใช้ความรู้และเทคโนโลยีใหม่ได้ด้วยตนเอง และสามารถทำงานเป็นทีม PLO20	วิเคราะห์และออกแบบระบบงานด้านเทคโนโลยีสารสนเทศได้ PLO21	วิเคราะห์ วางแผน หรือพัฒนาระบบงานที่มีการบูรณาการความรู้ในสาขาเทคโนโลยีสารสนเทศที่สามารถใช้งานได้',
    'https://www.cp.su.ac.th/file/show/1118'),
((SELECT degree_id FROM degree WHERE degree_name = 'ปริญญาตรี' LIMIT 1),
    (SELECT major_id FROM majors WHERE major_name = 'วิทยาการคอมพิวเตอร์' LIMIT 1),
    2565,'(วท.บ) หลักสูตรวิทยาศาสตรบัณฑิต สาขาวิชาวิทยาการคอมพิวเตอร์ 2565','Bachelor of Science Program in Computer Science 2022',
    'วิทยาศาสตรบัณฑิต (วิทยาการคอมพิวเตอร์)','Bachelor of Science (Computer Science)','สำเร็จการศึกษาระดับมัธยมศึกษาปีที่ 6 หรือเทียบเท่า','เกรดเฉลี่ยไม่ต่ำกว่า 2.00 เกรดเฉลี่ยวิชาเฉพาะไม่ต่ำกว่า 2.00',
    'ผลิตบัณฑิตผู้ผสานศาสตร์และศิลป์ที่พร้อมจะวางแผนก่อนปฏิบัติการ ทํางานมุ่งผลลัพธ์ ตรวจวัดผลสัมฤทธิ์ คิดวิเคราะห์เพื่อพัฒนาสร้างสรรค์สังคม',
    '1) เพื่อผลิตบัณฑิตที่มีความรู้เชิงทฤษฎีทางวิทยาการคอมพิวเตอร์และความสามารถเชิงปฏิบัติ พัฒนางานทางด้านคอมพิวเตอร์ได้อย่างมีประสิทธิภาพ หรือศึกษาต่อในระดับสูงขึ้น 2) เพื่อผลิตบัณฑิตที่มีทักษะสร้างสรรค์นวัตกรรมจากความรู้ด้านวิทยาการคอมพิวเตอร์ 3) เพื่อผลิตบัณฑิตที่สามารถถ่ายทอดความรู้ด้านวิทยาการคอมพิวเตอร์แก่ชุมชนและสังคม เพื่อการพัฒนาสังคมและประเทศชาติได้ 4) เพื่อผลิตบัณฑิตที่ปฏิบัติตนในกรอบจริยธรรมภายใต้กฎหมายเกี่ยวกับคอมพิวเตอร์และข้อมูล',
    'ประมาณ 20,000 บาทต่อเทอม','จำนวนไม่น้อยกว่า 126 หน่วยกิต',
    '1) โปรแกรมเมอร์ (Programmer) 2) นักวิทยาศาสตร์คอมพิวเตอร์ (Computer Scientist) 3) นักพัฒนาเว็บอย่างเต็มรูปแบบ (Full Stack Developer) 4) นักพัฒนาซอฟต์แวร์ (Software Developer) 5) นักวิเคราะห์และออกแบบระบบ (System Analyst and Designer) 6) นักวิทยาศาตร์ข้อมูล (Data Scientist) 7) ผู้ดูแลระบบเครือข่ายและเครื่องแม่ข่าย (Network and Server Administrator) 8) นักพัฒนาระบบอัตโนมัติของหุ่นยนต์ (Robotic Process Automation Developer)',
    'PLO1 อธิบายความหมายและคุณค่าของศิลปะและการสร้างสรรค์ได้ PLO2 อภิปรายความหมายของความหลากหลายทางวัฒนธรรมได้ PLO3 ระบุความรู้เบื้องต้นเกี่ยวกับการประกอบธุรกิจและทักษะพื้นฐานที่จำเป็นต่อการเป็นผู้ประกอบการได้ PLO4 มีทักษะการใช้ภาษา และสื่อสารได้ตรงตามวัตถุประสงค์ในบริบทการสื่อสารที่ หลากหลาย PLO5 เลือกใช้เทคโนโลยีสารสนเทศและการสื่อสารได้ตรงตามวัตถุประสงค์ ตลอดจนรู้เท่าทันสื่อและสารสนเทศ PLO6 แสวงหาความรู้ได้ด้วยตนเอง และนำความรู้ไปใช้ในการพัฒนา ตนเองและการดำเนินชีวิต PLO7 แสดงออกซึ่งทักษะความสัมพันธ์ระหว่างบุคคล สามารถทำงานร่วมกับผู้อื่นได้ มีระเบียบวินัย ตรงต่อเวลา ซื่อสัตย์สุจริต มีความรับผิดชอบต่อตนเอง สังคม และสิ่งแวดล้อม PLO8 ใช้ความคิดสร้างสรรค์ในการสร้างผลงานหรือดำเนินโครงการได้ PLO9 คิดวิเคราะห์ วางแผนอย่างเป็นระบบ เพื่อแก้ไขปัญหาหรือเพื่อออกแบบนวัตกรรมได้ PLO10 อธิบายหลักการทํางานและ แนวคิดของระบบและเทคโนโลยี ทางด้านคอมพิวเตอร์ สารสนเทศ และการสื่อสาร PLO11 จัดการระบบไฟล์ข้อมูลและระบบ ฐานข้อมูลตามบริบทของปัญหา PLO12 ประยุกต์ใช้อัลกอริทึมและโปรแกรมในการแก้ปัญหา ทางด้านคอมพิวเตอร์ภายใต้ สภาวะแวดล้อมที่กําหนด PLO13 กําหนดขอบเขตการทํางาน ติดตั้ง พร้อมตั้งค่าการใช้งานระบบ คอมพิวเตอร์และเครือข่าย PLO14 ประยุกต์ใช้ความรู้ในการ แก้ปัญหาด้านการรักษาความ มั่นคงปลอดภัยของข้อมูลของ ระบบคอมพิวเตอร์ภายใต้ กฎหมายโดยยึดหลักจริยธรรม ของการใช้ข้อมูลและสารสนเทศ PLO15 ปฏิบัติงานภายใต้รูปแบบของการ ทํางานเป็นทีมเพื่อบรรลุเป้าหมาย ในการดําเนินงานร่วมกัน PLO16 พูดและเขียนทั้งภาษาไทยและภาษาอังกฤษ เพื่อสื่อสารทําความ เข้าใจในด้านวิทยาการ คอมพิวเตอร์ PLO17 ติดตามข่าวสาร ข้อมูลและ ความก้าวหน้าทางเทคโนโลยีที่ เกี่ยวข้องกับงานด้านวิทยาการคอมพิวเตอร์ PLO18 วิเคราะห์ ออกแบบและพัฒนา ระบบคอมพิวเตอร์เพื่อแก้ไข ปัญหาให้ตรงตามความต้องการ ของผู้ใช้และบูรณาการตามบริบทของสังคม',
    'https://www.cp.su.ac.th/file/show/1119');

-- insert roadmap

INSERT INTO roadmap(course_id,roadmap_url) VALUES
((SELECT course_id FROM courses WHERE thai_course = '(วท.บ) หลักสูตรวิทยาศาสตรบัณฑิต สาขาวิชาเทคโนโลยีสารสนเทศ 2565' LIMIT 1),
    'https://cpsu-website.s3.ap-southeast-2.amazonaws.com/images/course/roadmap_IT_65.jpg'),
((SELECT course_id FROM courses WHERE thai_course = '(วท.บ) หลักสูตรวิทยาศาสตรบัณฑิต สาขาวิชาวิทยาการคอมพิวเตอร์ 2565' LIMIT 1),
    'https://cpsu-website.s3.ap-southeast-2.amazonaws.com/images/course/roadmap_CS_65.jpg');


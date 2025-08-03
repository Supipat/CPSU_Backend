-- CREATE news

CREATE TABLE IF NOT EXISTS news_type (
    type_id SERIAL PRIMARY KEY,
    type_name VARCHAR(255) NOT NULL
);

CREATE TABLE IF NOT EXISTS news (
    news_id SERIAL PRIMARY KEY,
    title VARCHAR(255) NOT NULL,
    content TEXT NOT NULL,
    type_id INT NOT NULL,
    detail_url TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (type_id) REFERENCES news_type(type_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS news_image (
    image_id SERIAL PRIMARY KEY,
    news_id INT NOT NULL,
    image_url TEXT NOT NULL,
    FOREIGN KEY (news_id) REFERENCES news(news_id) ON DELETE CASCADE
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

-- INSERT news

INSERT INTO news_type(type_name) VALUES
('ข่าวประชาสัมพันธ์'),
('ทุนการศึกษา'),
('รางวัลที่ได้รับ'),
('กิจกรรมของภาควิชา');

INSERT INTO news(title,content,type_id,detail_url) VALUES
('คู่มือแนะนำนักศึกษาใหม่ ปีการศึกษา 2568',
    'คู่มือแนะนำ การลงทะเบียนรายวิชาเรียน การเพิ่ม - ถอน - การเปลี่ยนกลุ่มเรียน ในรายวิชาเดิม การถอนรายวิชาโดยติดสัญลักษณ์ W ช่องทางการชำระค่าธรรมเนียมการศึกษา เอกสารแนบเบิกค่าเล่าเรียน การลาพักการศึกษา นักศึกษาที่คาดว่าจะสำเร็จการศึกษา การขอหนังสือสำคัญทางการศึกษา การติดต่อเจ้าหน้าที่งานทะเบียนและประมวลผล ช่องทางติดต่อคณะวิชา กองกิจการนักศึกษา การให้บริการสำหรับนักศึกษา วิทยาเขตสารสนเทศเพชรบุรี สำนักดิจิทัลเทคโนโลยี หอสมุด มหาวิทยาลัยศิลปากร ประกันอุบัติเหตุส่วนบุคคล ศูนย์บริการสุขภาพ คณะเภสัชศาสตร์ บริการด้านสุขภาพร่างกาย เครื่องแบบนักศึกษา / บัตรนักศึกษาอิเล็กทรอนิกส์',
    (SELECT type_id FROM news_type WHERE type_name = 'ข่าวประชาสัมพันธ์' LIMIT 1),
    'https://drive.google.com/file/d/1FnJRketlluku27HQqs-mVPnxMJ5wFxqZ/view?usp=sharing&fbclid=IwZXh0bgNhZW0CMTAAAR3M6TbS4tN1DUSa-z2NaM1ekAOELZp7TnsIhJWC6g_dvfz-sD_b0La0S7U_aem_K5_Zhi2eLKhIpJxaeszOlQ'),
('เปิดรับสมัครปริญญาโทและปริญญาเอกสาขาเทคโนโลยีสารสนเทศและนวัตกรรมดิจิทัล',
    'เปิดรับสมัครแล้ว! ภาควิชาคอมพิวเตอร์ คณะวิทยาศาสตร์ มหาวิทยาลัยศิลปากร เปิดรับสมัครนักศึกษาระดับบัณฑิตศึกษา ปริญญาโท - ปริญญาเอก หลักสูตร IT สำหรับคนทำงาน สู่ผู้เชี่ยวชาญด้านวิจัยและนวัตกรรม เน้นเรียนรู้กระบวนทำวิจัยเพื่อแก้ปัญหาจริง มีทุนผู้ช่วยวิจัยและทุนนำเสนอผลงานในงานประชุมวิชาการ มีทั้งภาคปกติ (เรียนวันธรรมดา) และโครงการพิเศษ (เรียน ส อา) มี 3 แผนการเรียนให้เลือก Data Science, Project Management, DevOps เปิดรับสมัครรอบ 1 วันที่ 16 ธ.ค. 67 - 21 ก.พ. 68 เปิดรับสมัครรอบ 2 วันี่ต่ 3 มี.ค. 68 - 7 พ.ค. 68 สมัครง่าย ๆ ผ่านช่องทางออนไลน์ได้ที่ https://graduate.su.ac.th มีข้อสงสัยสอบถามเพิ่มเติมได้ที่ เพจ Facebook : https://www.facebook.com/computingsu/ หรือ โทร 034-272-923',
    (SELECT type_id FROM news_type WHERE type_name = 'ข่าวประชาสัมพันธ์' LIMIT 1),
    'https://graduate.su.ac.th'),
('กำหนดการรับสมัครและสัมภาษณ์ทุนการศึกษา ประจำปีการศึกษา 2568 (ครั้งที่ 1)',
    'กำหนดการรับสมัครและสัมภาษณ์ทุนการศึกษา ของคณะวิทยาศาสตร์ เพื่อขอทุนการศึกษาสำหรับนักศึกษาปริญญาตรี ภาคการศึกษาต้น ปีการศึกษา 2568 สังกัดคณะวิทย์ ขาดแคลนทุน เกรดไม่ต่ำกว่า 2.00 ขอและยื่นใบสัมครได้ที่ " งานบริการการศึกษา ชั้น 1 อาคารวิทย์ 1 "',
    (SELECT type_id FROM news_type WHERE type_name = 'ทุนการศึกษา' LIMIT 1),
    'https://dsa.su.ac.th/ksu/wp-content/uploads/2025/06/%E0%B8%9B%E0%B8%A3%E0%B8%B0%E0%B8%81%E0%B8%B2%E0%B8%A8%E0%B8%81%E0%B8%AD%E0%B8%87%E0%B8%81%E0%B8%B4%E0%B8%88%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B8%99%E0%B8%B1%E0%B8%81%E0%B8%A8%E0%B8%B6%E0%B8%81%E0%B8%A9%E0%B8%B2-%E0%B9%80%E0%B8%A3%E0%B8%B7%E0%B9%88%E0%B8%AD%E0%B8%87-%E0%B8%81%E0%B8%B3%E0%B8%AB%E0%B8%99%E0%B8%94%E0%B8%81%E0%B8%B2%E0%B8%A3%E0%B8%A3.pdf'),
('ขอแสดงความยินดีกับ ผู้ช่วยศาสตราจารย์ ดร.อรวรรณ เชาวลิต อาจารย์ประจำภาควิชาคอมพิวเตอร์',
    'คณะวิทยาศาสตร์ มหาวิทยาลัยศิลปากร ขอแสดงความยินดีกับ ผู้ช่วยศาสตราจารย์ ดร.อรวรรณ เชาวลิต ภาควิชาคอมพิวเตอร์ ในโอกาสที่ตีพิมพ์ผลงานวิจัยในวารสาร ICIC Express Letters, Part B: Applications ในฐานข้อมูล Scopus (ScimagoJR, Quartile 4) เรื่อง Multivariate time series forecasting Thailand soybean meal price with deep learning models ICIC Express Letters, Part B: Applications 2025, 16(3), 317-323. http://www.icicelb.org/ellb/contents/2025/3/elb-16-03-10.pdf SDG 8 DECENT WORK AND ECONOMIC GROWTH',
    (SELECT type_id FROM news_type WHERE type_name = 'รางวัลที่ได้รับ' LIMIT 1),
    'http://www.icicelb.org/ellb/contents/2025/3/elb-16-03-10.pdf'),
('ขอแสดงความยินดีกับผู้ช่วยศาสตราจารย์ ดร.ปัญญนัท อ้นพงษ์ เนื่องในโอกาสได้รับการแต่งตั้งให้ดำรงตำแหน่งทางวิชาการ',
    'ภาควิชาคอมพิวเตอร์ คณะวิทยาศาสตร์ มหาวิทยาลัยศิลปากร ขอแสดงความยินดีกับผู้ช่วยศาสตราจารย์ ดร.ปัญญนัท อ้นพงษ์ เนื่องในโอกาสได้รับการแต่งตั้งให้ดำรงตำแหน่งทางวิชาการ ตำแหน่ง “ผู้ช่วยศาสตราจารย์” ในสาขาวิชาเทคโนโลยีสารสนเทศ (1806)',
    (SELECT type_id FROM news_type WHERE type_name = 'รางวัลที่ได้รับ' LIMIT 1),
    'https://www.facebook.com/ScienceSilpakornUniversity/posts/คณะวิทยาศาสตร์-มหาวิทยาลัยศิลปากร-ขอแสดงความยินดีกับผู้ช่วยศาสตราจารย์-ดรปัญญนัท/982826210555894/'),
('ขอแสดงความยินดีกับ นายภากร กัทชลี นักศึกษา ปริญญาเอก หลักสูตรเทคโนโลยีสารสนเทศและนวัตกรรมดิจิทัล',
    'ภาควิชาคอมพิวเตอร์ ขอแสดงความยินดีกับ นายภากร กัทชลี นักศึกษา ปริญญาเอก หลักสูตรเทคโนโลยีสารสนเทศและนวัตกรรมดิจิทัล ที่ได้รับรางวัล Best Presentation ในงาน The 6th Asia Joint Conference on Computing (AJCC 2025) ณ เมือง Osaka ประเทศญี่ปุ่นจาก Paper หัวข้อ "Intention Classification of Chinese Topics on Thai Facebook Pages Using Transformer Models with Emotional Features" อาจารย์ที่ปรึกษา อ.ดร.สัจจาภรณ์ ไวจรรยา และ ผศ.ดร.ณัฐโชติ พรหมฤทธิ์',
    (SELECT type_id FROM news_type WHERE type_name = 'รางวัลที่ได้รับ' LIMIT 1),
    'https://www.facebook.com/computingsu/posts/-ภาควิชาคอมพิวเตอร์-ขอแสดงความยินดีกับ-นายภากร-กัทชลี-นักศึกษา-ปริญญาเอก-หลักสูต/122207080106177331/'),
('ขอแสดงความยินดีกับ  นายณัฐพิสิษ กังวานธรรมกุล,นายนัทธพงศ์ เป็กทองและนาย พิรพัฒน์ ยิ่งแก้ว นักศึกษาสาขาวิชาเทคโนโลยีสารสนเทศ กับผลงาน "การพัฒนาระบบอํานวยการสอบ คณะวิทยาศาสตร์ มหาวิทยาลัยศิลปากร',
    'ภาควิชาคอมพิวเตอร์ ขอแสดงความยินดีกับนักศึกษาและอาจารย์ที่ปรึกษา ที่เข้าร่วมนำเสนอผลงานและได้รับรางวัลในงาน การประชุมวิชาการระดับปริญญาตรีด้านคอมพิวเตอร์ภูมิภาคเอเชีย ครั้งที่ 13 (The 13th Asia Undergraduate Conference on Computing: AUCC 2025) ในระหว่างวันที่ 19 - 21 กุมภาพันธ์ พ.ศ. 2568 ณ มหาวิทยาลัยราชภัฏอุตรดิตถ์ จังหวัดอุตรดิตถ์ รางวัลรองชนะเลิศอันดับ 1 Track นวัตกรรม กับผลงาน "การพัฒนาระบบอํานวยการสอบ คณะวิทยาศาสตร์ มหาวิทยาลัยศิลปากร (Development of examination Management system for the Faculty of Science)" ซึ่งผลงานนี้ได้มีการนำไปใช้จริงทดแทนระบบเดิม ในการจัดสอบของคณะฯ เมื่อช่วงสอบกลางภาคการศึกษาที่ผ่านมา นักศึกษาที่เข้านำเสนอ - นาย ณัฐพิสิษ กังวานธรรมกุล - นาย นัทธพงศ์ เป็กทอง - นาย พิรพัฒน์ ยิ่งแก้ว นักศึกษาสาขาวิชาเทคโนโลยีสารสนเทศ อาจารย์ที่ปรึกษา ผศ.ดร.ณัฐโชติ พรหมฤทธิ์ และ อ.ดร.สัจจาภรณ์ ไวจรรยา',
    (SELECT type_id FROM news_type WHERE type_name = 'รางวัลที่ได้รับ' LIMIT 1),
    'https://www.facebook.com/computingsu/posts/-ภาควิชาคอมพิวเตอร์-ขอแสดงความยินดีกับนักศึกษาและอาจารย์ที่ปรึกษา-ที่เข้าร่วมนำเ/122194936622177331/'),
('งานสานสัมพันธ์ภาคคอมพิวเตอร์',
    'เตรียมตัวให้พร้อม! งานสานสัมพันธ์ภาคคอมพิวเตอร์กำลังจะเริ่มขึ้น! เจอกันวันที่ 28 กุมภาพันธ์ 2568 ณ ลานจอดรถตึก 4 เวลา 17:30 - 21:30 น. ธีม: ย้อนยุค ไม่จำกัดช่วงเวลา! จะเป็นยุคหิน มนุษย์ถ้ำ อารยธรรมโบราณ หรือยุคใดก็ได้ จัดเต็มมาให้สุด! มีรางวัลแต่งกายยอดเยี่ยม 3 รางวัล!ใครแต่งตัวเข้าธีมและโดดเด่นที่สุด มีสิทธิ์คว้ารางวัลไปเลย! สนุกไปกับการแสดงสุดพิเศษ, เกมสุดมันส์, ลุ้นรางวัล และดนตรีสดปิดท้าย! ใครลงชื่อไว้แล้ว เจอกันแน่นอน! เตรียมชุดให้พร้อม แล้วมาสนุกไปด้วยกัน!',
    (SELECT type_id FROM news_type WHERE type_name = 'กิจกรรมของภาควิชา' LIMIT 1),
    '');

INSERT INTO news_image(news_id,image_url) VALUES
((SELECT news_id FROM news WHERE title = 'คู่มือแนะนำนักศึกษาใหม่ ปีการศึกษา 2568' LIMIT 1),
    'https://scontent.fbkk17-1.fna.fbcdn.net/v/t39.30808-6/476724160_122193063290177331_3666925435167000830_n.jpg?stp=dst-jpg_s600x600_tt6&_nc_cat=101&ccb=1-7&_nc_sid=127cfc&_nc_ohc=CFqfkUdW-dAQ7kNvwEoQJ7O&_nc_oc=AdnYcG432hVJbuWHUGruSwPFdYEYnbbEcw9VV9FbEAQNZJ_TfPEW5sU8vQ0w241Wefo&_nc_zt=23&_nc_ht=scontent.fbkk17-1.fna&_nc_gid=qEs6T26HiJ4Fa21Eiz0ZsQ&oh=00_AfRkDcleWsdLr56zBPsLld8Z1utvW9qRcLOTcngFgdFItQ&oe=6882B76A'),
((SELECT news_id FROM news WHERE title = 'เปิดรับสมัครปริญญาโทและปริญญาเอกสาขาเทคโนโลยีสารสนเทศและนวัตกรรมดิจิทัล' LIMIT 1),
    'https://scontent.fbkk17-1.fna.fbcdn.net/v/t39.30808-6/476889536_122192816792177331_4418576378512203331_n.jpg?stp=dst-jpg_s600x600_tt6&_nc_cat=101&ccb=1-7&_nc_sid=127cfc&_nc_ohc=j7sqipmw-PQQ7kNvwGtrMEk&_nc_oc=AdnrwECgbYOxhZ9kNTq7j_2MMZ8DEf3LtAYfb-u3lpMK9SAxOV_MG6dJ0bDvcbUpZoQ&_nc_zt=23&_nc_ht=scontent.fbkk17-1.fna&_nc_gid=oWQh2yHJoaPfbFrZm_WFDQ&oh=00_AfTo6HpbaSZb7raTRxlpvn0di_q2v-jsUou_ndgLwI9f7Q&oe=6882AD62'),
((SELECT news_id FROM news WHERE title = 'เปิดรับสมัครปริญญาโทและปริญญาเอกสาขาเทคโนโลยีสารสนเทศและนวัตกรรมดิจิทัล' LIMIT 1),
    'https://scontent.fbkk17-1.fna.fbcdn.net/v/t39.30808-6/476462814_122192816990177331_5763561513784129452_n.jpg?stp=dst-jpg_p417x417_tt6&_nc_cat=102&ccb=1-7&_nc_sid=127cfc&_nc_ohc=BVbfdczxOokQ7kNvwFd7VlT&_nc_oc=AdltGdACV7oehmAf64ZqN088CGr4Np-6KKHZct9m9RlLtKOHPsHjJX7KqMqT0ayMePo&_nc_zt=23&_nc_ht=scontent.fbkk17-1.fna&_nc_gid=oWQh2yHJoaPfbFrZm_WFDQ&oh=00_AfT575ZN1ZEIKq3vzGpvdxtZUjZjb1Deg8rO0-T56tZykw&oe=6882C2D2'),
((SELECT news_id FROM news WHERE title = 'เปิดรับสมัครปริญญาโทและปริญญาเอกสาขาเทคโนโลยีสารสนเทศและนวัตกรรมดิจิทัล' LIMIT 1),
    'https://scontent.fbkk17-1.fna.fbcdn.net/v/t39.30808-6/477254992_122192817320177331_1619181626657327777_n.jpg?stp=dst-jpg_p417x417_tt6&_nc_cat=107&ccb=1-7&_nc_sid=127cfc&_nc_ohc=Ds2QZ8ld0z4Q7kNvwGgIlpR&_nc_oc=AdmPKEKx5L6BhjLkiV-PwlaTC6PBpHz0oDnFUnFgyt21USYTDpfjL5JO_USTlWD5wGs&_nc_zt=23&_nc_ht=scontent.fbkk17-1.fna&_nc_gid=oWQh2yHJoaPfbFrZm_WFDQ&oh=00_AfQNwzPpjCqP-i1RgbQ2O57fiIX0VBxr87mX2ouiSoAafw&oe=6882A9A9'),
((SELECT news_id FROM news WHERE title = 'เปิดรับสมัครปริญญาโทและปริญญาเอกสาขาเทคโนโลยีสารสนเทศและนวัตกรรมดิจิทัล' LIMIT 1),
    'https://scontent.fbkk17-1.fna.fbcdn.net/v/t39.30808-6/476956494_122192817008177331_4161711359219247080_n.jpg?stp=dst-jpg_p417x417_tt6&_nc_cat=110&ccb=1-7&_nc_sid=127cfc&_nc_ohc=wDrJHZMPNiwQ7kNvwFYRA7d&_nc_oc=AdnW6jz1sHA5d8WpnyzkBK77_vLs1x-xrRa5UVg9894PNuPB5pw7aeIOIYxeVG7MCIE&_nc_zt=23&_nc_ht=scontent.fbkk17-1.fna&_nc_gid=oWQh2yHJoaPfbFrZm_WFDQ&oh=00_AfTtYlKG5nnWbbs_-gQD10eHdigI3Rf46mOdjCS0Z6EkfQ&oe=6882AFA7'),
((SELECT news_id FROM news WHERE title = 'กำหนดการรับสมัครและสัมภาษณ์ทุนการศึกษา ประจำปีการศึกษา 2568 (ครั้งที่ 1)' LIMIT 1),
    'https://scontent.fbkk17-1.fna.fbcdn.net/v/t51.75761-15/503974608_18043189844629140_2819006082379639761_n.jpg?stp=dst-jpg_p526x296_tt6&_nc_cat=108&ccb=1-7&_nc_sid=127cfc&_nc_ohc=lAyk2n25u9MQ7kNvwHJuHK2&_nc_oc=AdlYVvmyc9y3p9UBBWOKIjZp4Hj9O06RTgsljJa8Fv1uwR8GOZOpDQ4JWhR4WHh7j34&_nc_zt=23&_nc_ht=scontent.fbkk17-1.fna&_nc_gid=77Ksm8NPzMsMbIdxJsQSQw&oh=00_AfRfJbFMcVBbPj5G0aGt3-JvbFJP20n_ysAE-WB_2G7WPA&oe=6882AC14'),
((SELECT news_id FROM news WHERE title = 'ขอแสดงความยินดีกับ ผู้ช่วยศาสตราจารย์ ดร.อรวรรณ เชาวลิต อาจารย์ประจำภาควิชาคอมพิวเตอร์' LIMIT 1),
    'https://scontent.fbkk17-1.fna.fbcdn.net/v/t39.30808-6/500769500_1125022333002947_6103109281669185312_n.jpg?stp=dst-jpg_p526x296_tt6&_nc_cat=102&ccb=1-7&_nc_sid=f727a1&_nc_ohc=vJcBWdANXC4Q7kNvwGLmDXw&_nc_oc=Adka_y3kZ_zoLj-5XH33oqTLdnvvQ5C_4LmUV6Q88cKdoArgtKhOyo60kB8pO5TQOC4&_nc_zt=23&_nc_ht=scontent.fbkk17-1.fna&_nc_gid=M0tqf02zdIZFC9kCX4ABmQ&oh=00_AfQ3kDFRRPNX34uFFutIvdvX3bbqyLULBUhe-BSQo2GQdw&oe=68829F2D'),
((SELECT news_id FROM news WHERE title = 'ขอแสดงความยินดีกับผู้ช่วยศาสตราจารย์ ดร.ปัญญนัท อ้นพงษ์ เนื่องในโอกาสได้รับการแต่งตั้งให้ดำรงตำแหน่งทางวิชาการ' LIMIT 1),
    'https://scontent.fbkk17-1.fna.fbcdn.net/v/t39.30808-6/476354497_122192603522177331_8363469383456187557_n.jpg?stp=dst-jpg_p526x296_tt6&_nc_cat=104&ccb=1-7&_nc_sid=127cfc&_nc_ohc=kiW16bpKnzcQ7kNvwF1K2Yu&_nc_oc=AdmCo7Kzi0xWUgmEEK6NcV5Hj97Ha9kct3KaYFX4CUz_yGbjj11bFkW_NhvV5U3S0nY&_nc_zt=23&_nc_ht=scontent.fbkk17-1.fna&_nc_gid=lNRCvRZ1w_R8yBCF36iiLQ&oh=00_AfTq8wgPlhreg6BPlbqFVaizGyIWaAzUReg1CaP9igB3hw&oe=68829895'),
((SELECT news_id FROM news WHERE title = 'ขอแสดงความยินดีกับ นายภากร กัทชลี นักศึกษา ปริญญาเอก หลักสูตรเทคโนโลยีสารสนเทศและนวัตกรรมดิจิทัล' LIMIT 1),
    'https://scontent.fbkk17-1.fna.fbcdn.net/v/t39.30808-6/493228156_122207080010177331_3069128138840782487_n.jpg?stp=cp6_dst-jpg_s600x600_tt6&_nc_cat=107&ccb=1-7&_nc_sid=127cfc&_nc_ohc=C0by-1mKXO8Q7kNvwG4cc1W&_nc_oc=Adl7k2-g5ZaDBxwqK-6VmZBwF77McGl0iVl9KMCL1Ug9K3ohBoUcKJCi_xB3hoEoCAw&_nc_zt=23&_nc_ht=scontent.fbkk17-1.fna&_nc_gid=JXawUh5eC9aE2lQMcOuvPA&oh=00_AfSrq2xIlZ_9VYJWandWjZCdr_yDAJZaPaco7Esul0EdzA&oe=6882B08E'),
((SELECT news_id FROM news WHERE title = 'ขอแสดงความยินดีกับ นายภากร กัทชลี นักศึกษา ปริญญาเอก หลักสูตรเทคโนโลยีสารสนเทศและนวัตกรรมดิจิทัล' LIMIT 1),
    'https://scontent.fbkk17-1.fna.fbcdn.net/v/t39.30808-6/492087235_122207080034177331_8029941393914891284_n.jpg?stp=cp6_dst-jpg_s600x600_tt6&_nc_cat=110&ccb=1-7&_nc_sid=127cfc&_nc_ohc=VUnEpIseGMcQ7kNvwE6hCAh&_nc_oc=AdmTXNogP0K8Xrt5XZxdJp-amPOyiA9VXXgQl3IHcAKvjnV3-_iTx0KVMLx_I2IsvLA&_nc_zt=23&_nc_ht=scontent.fbkk17-1.fna&_nc_gid=JXawUh5eC9aE2lQMcOuvPA&oh=00_AfRCwbY_C31NPu55cQSbt6PU6R-JVAovmEoa84AbHxQvsQ&oe=6882B71A'),
((SELECT news_id FROM news WHERE title = 'ขอแสดงความยินดีกับ  นายณัฐพิสิษ กังวานธรรมกุล,นายนัทธพงศ์ เป็กทองและนาย พิรพัฒน์ ยิ่งแก้ว นักศึกษาสาขาวิชาเทคโนโลยีสารสนเทศ กับผลงาน "การพัฒนาระบบอํานวยการสอบ คณะวิทยาศาสตร์ มหาวิทยาลัยศิลปากร' LIMIT 1),
    'https://scontent.fbkk17-1.fna.fbcdn.net/v/t39.30808-6/480387411_122194936274177331_3448749609783995479_n.jpg?stp=dst-jpg_s600x600_tt6&_nc_cat=106&ccb=1-7&_nc_sid=127cfc&_nc_ohc=CE08oIkVcVcQ7kNvwF5uEqG&_nc_oc=AdnFDyhnaSenhTRNrDLJZ0YCR3DaMSacYV3OvLUMCxn7g29n_v2Atw2wBCjPn4o_kX8&_nc_zt=23&_nc_ht=scontent.fbkk17-1.fna&_nc_gid=UWS6wyKMSJ1jvjPrO9LtYA&oh=00_AfRGhwS1oWzYEL5jUcXVylG5AtvTze8tdhfW9rmo5WltoA&oe=68829C3C'),
((SELECT news_id FROM news WHERE title = 'ขอแสดงความยินดีกับ  นายณัฐพิสิษ กังวานธรรมกุล,นายนัทธพงศ์ เป็กทองและนาย พิรพัฒน์ ยิ่งแก้ว นักศึกษาสาขาวิชาเทคโนโลยีสารสนเทศ กับผลงาน "การพัฒนาระบบอํานวยการสอบ คณะวิทยาศาสตร์ มหาวิทยาลัยศิลปากร' LIMIT 1),
    'https://scontent.fbkk17-1.fna.fbcdn.net/v/t39.30808-6/480737512_122194936172177331_1928747647964741133_n.jpg?stp=cp6_dst-jpg_s600x600_tt6&_nc_cat=109&ccb=1-7&_nc_sid=127cfc&_nc_ohc=BG2tQutKEjIQ7kNvwGve2PR&_nc_oc=Adnm198cpu6fjdYgXajGZrFTZXLtDG1z3b5lKbtjnwZi3-Fa2lOnecCgcPbd4A4QLZY&_nc_zt=23&_nc_ht=scontent.fbkk17-1.fna&_nc_gid=UWS6wyKMSJ1jvjPrO9LtYA&oh=00_AfTes9RP5QJMZSUs4YQPt1jZBA5pt9iPsnPBFU1ztjt4Rg&oe=6882BC96'),
((SELECT news_id FROM news WHERE title = 'ขอแสดงความยินดีกับ  นายณัฐพิสิษ กังวานธรรมกุล,นายนัทธพงศ์ เป็กทองและนาย พิรพัฒน์ ยิ่งแก้ว นักศึกษาสาขาวิชาเทคโนโลยีสารสนเทศ กับผลงาน "การพัฒนาระบบอํานวยการสอบ คณะวิทยาศาสตร์ มหาวิทยาลัยศิลปากร' LIMIT 1),
    'https://scontent.fbkk17-1.fna.fbcdn.net/v/t39.30808-6/480515232_122194936196177331_6478205379636797392_n.jpg?stp=cp6_dst-jpg_s600x600_tt6&_nc_cat=108&ccb=1-7&_nc_sid=127cfc&_nc_ohc=sw-NA63bHlcQ7kNvwGG8C3e&_nc_oc=Adn7l_suXANHSmTIGIP7Jhs-UVAjo_NzgUnpBUWjWAyLDmpDSDKz3R5Q44aX4eqOz5k&_nc_zt=23&_nc_ht=scontent.fbkk17-1.fna&_nc_gid=UWS6wyKMSJ1jvjPrO9LtYA&oh=00_AfRje8HfFDYbYK52MzwmG0NhH4hwv2VATqbm1Ddq4PH5lA&oe=6882BE8D'),
((SELECT news_id FROM news WHERE title = 'ขอแสดงความยินดีกับ  นายณัฐพิสิษ กังวานธรรมกุล,นายนัทธพงศ์ เป็กทองและนาย พิรพัฒน์ ยิ่งแก้ว นักศึกษาสาขาวิชาเทคโนโลยีสารสนเทศ กับผลงาน "การพัฒนาระบบอํานวยการสอบ คณะวิทยาศาสตร์ มหาวิทยาลัยศิลปากร' LIMIT 1),
    'https://scontent.fbkk17-1.fna.fbcdn.net/v/t39.30808-6/480781568_122194936526177331_3578900039569325654_n.jpg?stp=dst-jpg_s600x600_tt6&_nc_cat=108&ccb=1-7&_nc_sid=127cfc&_nc_ohc=LSYP9yq_dyAQ7kNvwElDCs0&_nc_oc=Adl67qCWtbKInJZdatfDSnvEORB3zv5ROeIjOJJH7MV41zgVHhr0TvutdNCTDwrJevs&_nc_zt=23&_nc_ht=scontent.fbkk17-1.fna&_nc_gid=UWS6wyKMSJ1jvjPrO9LtYA&oh=00_AfQqoFGlwYYFn3Dc5l_ykPgoqEXmaYXOtunmAP49zXbOBQ&oe=6882AEC8');

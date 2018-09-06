package cn.fantasticmao.util.moment;

import java.io.File;
import java.io.FileInputStream;
import java.io.IOException;

/**
 * FileUtil
 *
 * @author maodh
 * @since 2018/9/4
 */
class FileUtil {

    // common file utils

    static String getFileNameWithoutExtension(File file) {
        String fileName = file.getName();
        int lastIndexOf = fileName.lastIndexOf(".");
        return lastIndexOf == -1 ? "" : fileName.substring(0, lastIndexOf);
    }

    static String getFileExtension(File file) {
        String fileName = file.getName();
        int lastIndexOf = fileName.lastIndexOf(".");
        return lastIndexOf == -1 ? "" : fileName.substring(lastIndexOf);
    }

    static String getFileExtensionWithoutDot(File file) {
        String fileName = file.getName();
        int lastIndexOf = fileName.lastIndexOf(".");
        return lastIndexOf == -1 ? "" : fileName.substring(lastIndexOf + 1);
    }

    // image file utils

    /**
     * 比较图片的后缀名和魔数是否相符
     *
     * @see cn.fantasticmao.util.moment.ImageType
     */
    static boolean compareImageExtensionAndMagicNumber(File image) {
        // 图片的后缀名
        final String imageExtension = FileUtil.getFileExtension(image);
        // 图片的类型
        final ImageType imageType = ImageType.ofExtension(imageExtension);
        // 图片的魔数
        final String imageMagicNumber = FileUtil.getImageMagicNumber(image);
        return imageType != null && imageMagicNumber.startsWith(imageType.magicNumber);
    }

    private static String getImageMagicNumber(File file) {
        final int byteSize = 1 << 2;
        byte[] bytes = new byte[byteSize];
        try (FileInputStream in = new FileInputStream(file)) {
            in.read(bytes, 0, byteSize);
            return bytesToHex(bytes);
        } catch (IOException e) {
            e.printStackTrace();
            return "";
        }
    }

    private static final char[] HEX_ARRAY = "0123456789ABCDEF".toCharArray();

    private static String bytesToHex(byte[] bytes) {
        char[] hexChars = new char[bytes.length * 2];
        for (int j = 0; j < bytes.length; j++) {
            int v = bytes[j] & 0xFF;
            hexChars[j * 2] = HEX_ARRAY[v >>> 4];
            hexChars[j * 2 + 1] = HEX_ARRAY[v & 0x0F];
        }
        return new String(hexChars);
    }
}

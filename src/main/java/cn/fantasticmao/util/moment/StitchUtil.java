package cn.fantasticmao.util.moment;

import javax.imageio.ImageIO;
import java.awt.image.BufferedImage;
import java.io.File;
import java.io.IOException;
import java.util.List;
import java.util.Objects;

/**
 * StitchUtil
 *
 * @author maodh
 * @since 2018/9/4
 */
class StitchUtil {

    /**
     * @param imageFileList  图片文件列表
     * @param subtitleHeight 字幕高度，单位 px
     * @param outputDir      保存最终图片文件的目录
     * @return 拼接而成的最终图片文件
     * @throws IOException I/O Exception
     */
    static File stitchImageList(List<File> imageFileList, final int subtitleHeight, final String outputDir) throws IOException {
        // 校验参数
        if (imageFileList == null || imageFileList.size() == 0) {
            throw new IllegalArgumentException("the image list should not be empty");
        } else {
            if (imageFileList.stream().anyMatch(file -> !file.exists())) {
                throw new IllegalArgumentException("every list should be filled with image");
            }
        }
        if (subtitleHeight <= 0) {
            throw new IllegalArgumentException("the subtitle height should be a positive integer");
        }
        if (outputDir == null || outputDir.length() == 0) {
            throw new IllegalArgumentException("the output directory path should not be empty");
        } else {
            File outputFile = new File(outputDir);
            if (!outputFile.exists() || !outputFile.isDirectory()) {
                throw new IllegalArgumentException("there must be an output directory for finished image");
            }
        }

        final File firstImageFile = imageFileList.get(0);
        // 校验第一张图片文件的后缀名和魔数
        if (!FileUtil.compareImageExtensionAndMagicNumber(firstImageFile)) {
            throw new IllegalArgumentException("the format of No.1 image is not supported or cannot opened");
        }

        // 第一张图片的后缀名
        final String firstImageFileExtensionWithoutDot = FileUtil.getFileExtensionWithoutDot(firstImageFile);

        // 最终拼接而成的图片
        BufferedImage firstImage = ImageIO.read(firstImageFile);
        final int finalImageWidth = firstImage.getWidth();
        final int finalImageHeight = firstImage.getHeight() + (imageFileList.size() - 1) * subtitleHeight;
        final int finalImageType = firstImage.getType();
        BufferedImage finalImage = new BufferedImage(finalImageWidth, finalImageHeight, finalImageType);

        // 拼接图片
        finalImage.createGraphics().drawImage(firstImage, 0, 0, null);
        for (int i = 0; i < imageFileList.size(); i++) {
            if (i == 0) continue;
            File thisImageFile = imageFileList.get(i);

            // 校验当前图片文件的后缀名和魔数
            if (!FileUtil.compareImageExtensionAndMagicNumber(thisImageFile)) {
                throw new IllegalArgumentException("the format of No." + (i + 1) + " image is not supported or cannot opened");
            }

            // 校验当前图片文件的后缀名
            final String thisImageFileExtensionWithoutDot = FileUtil.getFileExtensionWithoutDot(thisImageFile);
            if (!Objects.equals(thisImageFileExtensionWithoutDot, firstImageFileExtensionWithoutDot)) {
                throw new IllegalArgumentException("the extension of the No." + (i + 1) + " image should be same as the first image");
            }

            BufferedImage thisImage = ImageIO.read(thisImageFile);

            // 校验当前图片的宽度
            if (thisImage.getWidth() != firstImage.getWidth()) {
                throw new IllegalArgumentException("the width of the No." + (i + 1) + " image should be same as the first image");
            }

            // 裁剪图片字幕
            BufferedImage subtitle = thisImage.getSubimage(0, thisImage.getHeight() - subtitleHeight, thisImage.getWidth(), subtitleHeight);

            // 拼接图片字幕
            finalImage.createGraphics().drawImage(subtitle, 0, firstImage.getHeight() + (i - 1) * subtitleHeight, null);
        }

        // 如果最终图片文件已存在，则会覆盖
        File finalImageFile = newFinalImageFileFromImageList(imageFileList, outputDir);
        ImageIO.write(finalImage, firstImageFileExtensionWithoutDot, finalImageFile);
        return finalImageFile;
    }

    private static File newFinalImageFileFromImageList(List<File> imageFileList, String outputDir) {
        final File firstImageFile = imageFileList.get(0);
        final String finalImageFileNameWithoutExtension = FileUtil.getFileNameWithoutExtension(firstImageFile);
        final String finalImageFileExtension = FileUtil.getFileExtension(firstImageFile);

        final String finalImagePath = outputDir + finalImageFileNameWithoutExtension + "_final" + finalImageFileExtension;
        return new File(finalImagePath);
    }

}
